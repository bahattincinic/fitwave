package database

import (
	"github.com/bahattincinic/fitwave/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertGears(tx *gorm.DB, gears []models.Gear) error {
	for _, row := range gears {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&row).Error
		if err != nil {
			return err
		}
	}

	return nil
}
