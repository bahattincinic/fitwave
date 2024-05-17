package database

import (
	"github.com/bahattincinic/fitwave/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertAthletes(tx *gorm.DB, athletes []models.Athlete) error {
	for _, row := range athletes {
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
