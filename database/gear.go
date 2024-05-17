package database

import (
	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertGears(tx *gorm.DB, gears []models.Gear) error {
	for _, row := range gears {
		currentRow := row
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&currentRow).Error
		if err != nil {
			d.log.Error("could not upsert gear",
				zap.Any("gear", currentRow))
			return err
		}
	}
	return nil
}

func (d *Database) ListGears(offset, limit int) (int64, []models.Gear, error) {
	var gears []models.Gear
	var count int64

	err := d.db.
		Limit(limit).
		Offset(offset).
		Preload("Athlete").
		Order("id desc").
		Find(&gears).
		Count(&count).
		Error

	if err != nil {
		return 0, nil, pkgerrors.New("error while fetching gears")
	}

	return count, gears, nil
}

func (d *Database) GetGear(id string) (*models.Gear, error) {
	var gear models.Gear
	if err := d.db.Preload("Athlete").First(&gear, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &gear, nil
}
