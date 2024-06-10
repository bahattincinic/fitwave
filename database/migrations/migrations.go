package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Migration struct {
	migrate *gormigrate.Gormigrate
}

func NewMigration(db *gorm.DB) *Migration {
	m := &Migration{}
	m.migrate = gormigrate.New(db, gormigrate.DefaultOptions, m.Get())
	return m
}

// Get returns the list of all migrations to be applied.
func (m *Migration) Get() []*gormigrate.Migration {
	var migrations []*gormigrate.Migration

	migrations = append(migrations, m.createStravaTables()...)
	migrations = append(migrations, m.createConfigTable()...)
	migrations = append(migrations, m.createDashboardTables()...)

	return migrations
}

// Migrate applies the migrations to the database.
func (m *Migration) Migrate() error {
	return m.migrate.Migrate()
}
