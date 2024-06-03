package config

import (
	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

var (
	// Env global environment accessible from anywhere
	Env string
)

const (
	EnvProduction = "production"
	EnvLocal      = "local"
)

type Config struct {
	Env      string `env:"ENV" envDefault:"local"`
	Database Database
	Log      LogConfig
	API      API
}

type LogConfig struct {
	Level  string `env:"LOG_LEVEL" envDefault:"debug"`
	Output string `env:"LOG_OUTPUT" envDefault:"stdout"`
}

type ConnectionType string

const (
	MySQL    ConnectionType = "mysql"
	Postgres ConnectionType = "postgresql"
	SQLite   ConnectionType = "sqlite"
)

type Database struct {
	DSN         string         `env:"DATABASE_DSN" envDefault:"fitWave.db"`
	Type        ConnectionType `env:"DATABASE_TYPE" envDefault:"sqlite"`
	AutoMigrate bool           `env:"DATABASE_AUTO_MIGRATE" envDefault:"true"`
}

type API struct {
	PORT            int    `env:"API_PORT" envDefault:"9000"`
	SecretKey       string `env:"API_SECRET_KEY" envDefault:"PBU_gha4zfk!rwj8axv"`
	TokenExpiryHour int    `env:"API_TOKEN_EXPIRY_HOUR" envDefault:"24"`
}

func Production() bool {
	return Env == EnvProduction
}

func Local() bool {
	return Env == EnvLocal
}

func Parse() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	Env = cfg.Env

	return &cfg, nil
}
