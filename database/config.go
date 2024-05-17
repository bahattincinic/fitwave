package database

import (
	"github.com/bahattincinic/fitwave/models"
	"gorm.io/gorm"
)

func (d *Database) GetCurrentConfig() (*models.Config, error) {
	var cfg models.Config

	err := d.db.
		Find(&cfg).
		Order("id desc").
		Limit(1).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Config{}, nil
		}
		return nil, err
	}

	return &cfg, err
}

func (d *Database) UpsertConfig(in models.Config) (*models.Config, error) {
	cfg, err := d.GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	cfg.AccessToken = in.AccessToken
	cfg.AthleteId = in.AthleteId
	cfg.ClientId = in.ClientId
	cfg.ClientSecret = in.ClientSecret

	if err := d.db.Save(&cfg).Error; err != nil {
		return nil, err
	}
	return cfg, nil
}
