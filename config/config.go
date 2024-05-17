package config

import (
	"encoding/json"
	"os"
	"path/filepath"
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
	Env      string       `json:"env"`
	Database Database     `json:"database"`
	Log      LogConfig    `json:"log"`
	Strava   StravaConfig `json:"strava"`
}

type LogConfig struct {
	Output string `json:"output"`
}

type StravaConfig struct {
	AccessToken string `json:"access_token"`
	AthleteId   int64  `json:"athlete_id"`
}

type Database struct {
	DSN         string `json:"dsn"`
	AutoMigrate bool   `json:"auto_migrate"`
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

func Parse(path string) (*Config, error) {
	c := Config{}

	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	Env = c.Env

	return &c, nil
}
