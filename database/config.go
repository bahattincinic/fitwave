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
