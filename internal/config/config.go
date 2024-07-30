package config

import (
	"fmt"
	"time"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type PensionsConfig struct {
	URL              string        `mapstructure:"PENSIONS_URL" default:"https://api7.pensions.ge/api"`
	AuthToken        string        `mapstructure:"PENSIONS_AUTH_TOKEN"`
	Username         string        `mapstructure:"PENSIONS_USERNAME"`
	Password         string        `mapstructure:"PENSIONS_PASSWORD"`
	ClientTimeout    time.Duration `mapstructure:"PENSIONS_CLIENT_TIMEOUT" default:"10s"`
	ClientRetryCount int           `mapstructure:"PENSIONS_CLIENT_RETRY_COUNT" default:"2"`
	ClientDebug      bool          `mapstructure:"PENSIONS_CLIENT_DEBUG" default:"false"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

type Config struct {
	Env string `mapstructure:"ENV" default:"local"`

	Pensions PensionsConfig `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
}

func New() (*Config, error) {
	conf := new(Config)

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	defaults.SetDefaults(conf)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return conf, nil
}
