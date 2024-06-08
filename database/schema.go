package database

import (
	"sync"

	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm/schema"
)

// Schema represents the schema information of a database field.
type Schema struct {
	TableName string `json:"table_name"`
	DBName    string `json:"field_db_name"`
	Type      string `json:"type"`
}

// GetModelsSchema retrieves the schema information of the database models.
// It parses the fields of predefined models (Activity, Athlete, Gear) and returns their schema details.
func (d *Database) GetModelsSchema() ([]Schema, error) {
	var resp []Schema

	dbModels := []interface{}{
		&models.Activity{},
		&models.Athlete{},
		&models.Gear{},
	}

	for _, model := range dbModels {
		s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			return nil, pkgerrors.Wrap(err, "Parse")
		}

		for _, field := range s.Fields {
			if field.DBName == "" {
				continue
			}
			resp = append(resp, Schema{
				TableName: s.Table,
				DBName:    field.DBName,
				Type:      string(field.DataType),
			})
		}
	}

	return resp, nil
}
