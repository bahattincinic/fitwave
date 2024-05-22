package database

import (
	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertActivities(tx *gorm.DB, activities []models.Activity) error {
	for _, row := range activities {
		currentRow := row
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&currentRow).Error
		if err != nil {
			d.log.Error("could not upsert activity",
				zap.Any("activity", currentRow))
			return err
		}
	}
	return nil
}

func (d *Database) GetLatestActivity() (*models.Activity, error) {
	var act models.Activity

	err := d.db.
		Find(&act).
		Order("id desc").
		Limit(1).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &act, err
}

func (d *Database) ListActivities(offset, limit int) (int64, []models.Activity, error) {
	var activities []models.Activity
	var count int64

	err := d.db.
		Model(&models.Activity{}).
		Count(&count).
		Error

	if err != nil {
		return 0, nil, pkgerrors.New("error while counting activities")
	}

	err = d.db.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Preload("Athlete").
		Preload("Gear").
		Find(&activities).
		Error

	if err != nil {
		return 0, nil, pkgerrors.New("error while fetching activities")
	}

	return count, activities, nil
}

func (d *Database) GetActivity(id string) (*models.Activity, error) {
	var act models.Activity

	err := d.db.
		Preload("Athlete").
		Preload("Gear").
		First(&act, id).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &act, nil
}
