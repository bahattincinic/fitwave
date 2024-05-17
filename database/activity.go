package database

import (
	"github.com/bahattincinic/fitwave/models"
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

	return &act, err
}
