package models

import (
	"time"
)

type Config struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	ClientId     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
