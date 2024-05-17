package database

import (
	"github.com/bahattincinic/fitwave/models"
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
