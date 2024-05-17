package database

import (
	"github.com/bahattincinic/fitwave/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpsertAthletes(tx *gorm.DB, athletes []models.Athlete) error {
	for _, row := range athletes {
		currentRow := row
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&currentRow).Error
		if err != nil {
			d.log.Error("could not upsert athletes",
				zap.Any("athlete", currentRow))
			return err
		}
	}
	return nil
}
