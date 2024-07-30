package processor

import (
	"fmt"
	"log/slog"

	"github.com/abgeo/pensions/internal/client"
	"github.com/abgeo/pensions/internal/config"
	"github.com/abgeo/pensions/internal/database"
	"github.com/abgeo/pensions/internal/model"
	"github.com/abgeo/pensions/internal/repository"
	"github.com/abgeo/pensions/internal/service"
	"github.com/abgeo/pensions/internal/service/auth"
	"github.com/abgeo/pensions/internal/service/contributions"
	"github.com/jinzhu/copier"
)

type Processor struct {
	logger            *slog.Logger
	config            *config.Config
	db                *database.Database
	contributionsRepo *repository.ContributionRepository
	authSvc           *auth.Auth
	contributionsSvc  *contributions.Contributions
	accessToken       string
}

func New() (*Processor, error) {
	var err error

	proc := &Processor{
		logger: slog.Default(),
	}

	proc.config, err = config.New()
	if err != nil {
		return nil, fmt.Errorf("unable to get config: %w", err)
	}

	proc.db, err = database.New(proc.config.Database)
	if err != nil {
		return nil, fmt.Errorf("unable to create database connection: %w", err)
	}

	proc.contributionsRepo, err = repository.NewContributionRepository(proc.db)
	if err != nil {
		return nil, fmt.Errorf("unable to create contributions repository: %w", err)
	}

	httpClient := client.New(proc.config.Pensions)
	proc.authSvc = auth.New(httpClient)
	proc.contributionsSvc = contributions.New(httpClient)

	return proc, nil
}

func (proc *Processor) Process() int {
	var err error

	proc.logger.Info("starting processor")

	proc.accessToken, err = proc.getAuthToken()
	if err != nil {
		proc.logger.Error("unable to get Access Token", slog.Any("error", err))

		return 1
	}

	if err = proc.processContributions(); err != nil {
		proc.logger.Error("unable to process contributions", slog.Any("error", err))

		return 1
	}

	proc.logger.Info("processing has been finished")

	return 0
}

func (proc *Processor) getAuthToken() (string, error) {
	if proc.config.Pensions.AuthToken != "" {
		return proc.config.Pensions.AuthToken, nil
	}

	proc.logger.Info("requesting new Access Token")

	data, _, err := proc.authSvc.Authenticate(proc.config.Pensions.Username, proc.config.Pensions.Password)
	if err != nil {
		return "", fmt.Errorf("unable to authenticate user: %w", err)
	}

	return data.AccessToken, nil
}

func (proc *Processor) processContributions() error {
	entity := &model.Contribution{}

	proc.logger.Info("fetching contributions")

	data, _, err := proc.contributionsSvc.Get(service.WithAuthToken(proc.accessToken))
	if err != nil {
		return fmt.Errorf("unable to fetch contributions: %w", err)
	}

	proc.logger.Info("contributions received", slog.String("userID", data.UserID.String()))

	_ = copier.Copy(&entity, &data)

	if err = proc.contributionsRepo.Create(entity); err != nil {
		return fmt.Errorf("unable to store contributions: %w", err)
	}

	return nil
}
