package database

import (
	"context"
	"fmt"

	"github.com/bahattincinic/fitwave/config"
	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	ctx context.Context
	db  *gorm.DB
	cfg *config.Config
	log *zap.Logger
}

func NewDatabase(ctx context.Context, log *zap.Logger, cfg *config.Config) (*Database, error) {
	var conn gorm.Dialector
	switch cfg.Database.Type {
	case config.SQLite:
		conn = sqlite.Open(cfg.Database.DSN)
	case config.Postgres:
		conn = postgres.Open(cfg.Database.DSN)
	case config.MySQL:
		conn = mysql.Open(cfg.Database.DSN)
	default:
		return nil, fmt.Errorf("invalid connection type: %s", cfg.Database.Type)
	}

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Open")
	}

	d := &Database{
		db:  db,
		ctx: ctx,
		log: log,
		cfg: cfg,
	}

	if cfg.Database.AutoMigrate {
		if err := d.Migrate(); err != nil {
			return nil, pkgerrors.Wrap(err, "Migrate")
		}
	}

	return d, nil
}

func (d *Database) Migrate() error {
	m := []interface{}{
		&models.Activity{},
		&models.Athlete{},
		&models.Gear{},
		&models.Config{},
	}
	return d.db.AutoMigrate(m...)
}

func (d *Database) BeginTransaction() *gorm.DB {
	return d.db.Begin()
}

func (d *Database) GetConnection() *gorm.DB {
	return d.db
}
