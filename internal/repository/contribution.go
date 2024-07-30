package repository

import (
	"fmt"

	"github.com/abgeo/pensions/internal/database"
	"github.com/abgeo/pensions/internal/model"
)

type ContributionRepository struct {
	db *database.Database
}

func NewContributionRepository(db *database.Database) (*ContributionRepository, error) {
	if err := db.AutoMigrate(&model.Contribution{}); err != nil {
		return nil, fmt.Errorf("unable to execute migration: %w", err)
	}

	return &ContributionRepository{
		db: db,
	}, nil
}

func (repo *ContributionRepository) Create(entity *model.Contribution) error {
	return repo.db.
		Create(entity).
		Error
}
