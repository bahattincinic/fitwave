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

type ConnectionType string

const (
	Mysql    ConnectionType = "mysql"
	Postgres ConnectionType = "postgresql"
	SQLITE   ConnectionType = "sqlite"
)

type Database struct {
	ctx context.Context
	db  *gorm.DB
	cfg *config.Config
	log *zap.Logger
}

func NewDatabase(ctx context.Context, log *zap.Logger, cfg *config.Config) (*Database, error) {
	var conn gorm.Dialector
	switch ConnectionType(cfg.Database.Type) {
	case SQLITE:
		conn = sqlite.Open(cfg.Database.DSN)
	case Postgres:
		conn = postgres.Open(cfg.Database.DSN)
	case Mysql:
		conn = mysql.Open(cfg.Database.DSN)
	default:
		return nil, fmt.Errorf("invalid connection type: %s", cfg.Database.Type)
	}

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Open")
	}

	{
		if cfg.Database.AutoMigrate {
			m := []interface{}{
				&models.Activity{},
				&models.Athlete{},
				&models.Gear{},
			}
			if err := db.AutoMigrate(m...); err != nil {
				return nil, pkgerrors.Wrap(err, "AutoMigrate")
			}
		}
	}

	return &Database{
		db:  db,
		ctx: ctx,
		log: log,
		cfg: cfg,
	}, nil
}

func (d *Database) BeginTransaction() *gorm.DB {
	return d.db.Begin()
}

func (d *Database) GetConnection() *gorm.DB {
	return d.db
}
