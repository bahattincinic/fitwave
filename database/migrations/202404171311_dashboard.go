package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// createDashboardTables returns the migrations for creating Dashboard-related tables.
func (m *Migration) createDashboardTables() []*gormigrate.Migration {
	return []*gormigrate.Migration{{
		ID: "202404171311",
		Migrate: func(tx *gorm.DB) error {
			type dashboard struct {
				ID        uint `gorm:"primarykey"`
				CreatedAt time.Time
				UpdatedAt time.Time
				Name      string
			}
			type component struct {
				ID          uint `gorm:"primarykey"`
				CreatedAt   time.Time
				UpdatedAt   time.Time
				DashboardID uint
				Dashboard   dashboard `gorm:"foreignkey:DashboardID"`
				Name        string
				Query       string
				Type        string
				Configs     datatypes.JSON
			}
			return tx.AutoMigrate(&dashboard{}, &component{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("dashboards", "components")
		},
	}}
}
