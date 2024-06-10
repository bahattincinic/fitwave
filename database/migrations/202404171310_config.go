package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// createConfigTable returns the migration for creating the Config table.
func (m *Migration) createConfigTable() []*gormigrate.Migration {
	return []*gormigrate.Migration{{
		ID: "202404171310",
		Migrate: func(tx *gorm.DB) error {
			type config struct {
				ID            uint `gorm:"primarykey"`
				CreatedAt     time.Time
				UpdatedAt     time.Time
				ClientId      int
				ClientSecret  string
				LoginType     string
				LoginUsername string
				LoginPassword string
			}
			return tx.AutoMigrate(&config{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("configs")
		},
	}}
}
