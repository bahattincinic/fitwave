package database

import (
	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertAthlete(tx *gorm.DB, athlete *models.Athlete) error {
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&athlete).Error
	if err != nil {
		d.log.Error("could not upsert athletes",
			zap.Any("athlete", athlete))
		return err
	}
	return nil
}

func (d *Database) ListAthletes(offset, limit int) (int64, []models.Athlete, error) {
	var athletes []models.Athlete
	var count int64

	err := d.db.
		Model(&models.Athlete{}).
		Count(&count).
		Error

	if err != nil {
		return 0, nil, pkgerrors.New("error while counting athletes")
	}

	err = d.db.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Find(&athletes).
		Error

	if err != nil {
		return 0, nil, pkgerrors.New("error while fetching athletes")
	}

	return count, athletes, nil
}

func (d *Database) GetAthlete(id string) (*models.Athlete, error) {
	var ath models.Athlete
	if err := d.db.Preload("Athlete").First(&ath, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ath, nil
}
