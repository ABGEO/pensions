package database

import (
	"fmt"
	"net"

	"github.com/abgeo/pensions/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(conf config.DatabaseConfig) (*Database, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		conf.User,
		conf.Password,
		net.JoinHostPort(conf.Host, conf.Port),
		conf.Database,
	)

	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &Database{
		DB: db,
	}, nil
}
