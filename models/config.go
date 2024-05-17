package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model

	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	AthleteId    int64  `json:"athlete_id"`
}
