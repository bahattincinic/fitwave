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
	Strava   StravaConfig
}

type LogConfig struct {
	Level  string `env:"LOG_LEVEL" envDefault:"debug"`
	Output string `env:"OUTPUT" envDefault:"stdout"`
}

type StravaConfig struct {
	AccessToken string `env:"STRAVA_ACCESS_TOKEN"`
	AthleteId   int64  `env:"STRAVA_ATHLETE_ID"`
}

type Database struct {
	DSN         string `env:"DATABASE_DSB"`
	Type        string `env:"DATABASE_TYPE"`
	AutoMigrate bool   `env:"DATABASE_AUTO_MIGRATE"`
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
