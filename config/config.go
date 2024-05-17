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
	EnvTesting    = "testing"
	EnvProduction = "production"
	EnvLocal      = "local"
)

type Config struct {
	Env      string `env:"ENV,notEmpty"`
	Database Database
	Log      LogConfig
	API      API
}

type LogConfig struct {
	Level  string `env:"LOG_LEVEL" envDefault:"debug"`
	Output string `env:"OUTPUT" envDefault:"stdout"`
}

type Database struct {
	DSN         string `env:"DATABASE_DSB"`
	Type        string `env:"DATABASE_TYPE"`
	AutoMigrate bool   `env:"DATABASE_AUTO_MIGRATE"`
}

type API struct {
	PORT int `env:"API_PORT" envDefault:"9000"`
}

func Production() bool {
	return Env == EnvProduction
}

func Testing() bool {
	return Env == EnvTesting
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
