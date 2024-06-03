package models

import (
	"time"
)

type LoginType string

const (
	AnonymousLoginType = LoginType("anonymous")
	ProtectedLoginType = LoginType("protected")
)

type Config struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	ClientId     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	LoginType     LoginType `json:"-"`
	LoginUsername string    `json:"-"`
	LoginPassword string    `json:"-"`
}

func (c *Config) SetupCompleted() bool {
	return c.ID != 0
}
