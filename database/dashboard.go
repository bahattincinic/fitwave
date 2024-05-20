package database

import (
	"encoding/json"

	"github.com/bahattincinic/fitwave/models"
	"gorm.io/gorm"
)

func (d *Database) ListDashboards(offset, limit int) (int64, []models.Dashboard, error) {
	var dashboards []models.Dashboard
	var count int64

	err := d.db.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Find(&dashboards).
		Count(&count).
		Error

	if err != nil {
		return 0, nil, err
	}

	return count, dashboards, nil
}

func (d *Database) GetDashboard(dashboardID uint) (*models.Dashboard, error) {
	var dashboard models.Dashboard

	if err := d.db.First(&dashboard, dashboardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &dashboard, nil
}

func (d *Database) ListComponents(dashboardId uint) ([]models.Component, error) {
	var components []models.Component
	err := d.db.
		Preload("Dashboard").
		Where("dashboard_id = ?", dashboardId).
		Find(&components).
		Error

	if err != nil {
		return nil, err
	}

	return components, nil
}

func (d *Database) DeleteDashboard(dashboard *models.Dashboard) error {
	r := d.db.Delete(&dashboard)

	return r.Error
}

func (d *Database) DeleteComponent(component *models.Component) error {
	r := d.db.Delete(&component)

	return r.Error
}

func (d *Database) CreateDashboard(name string) (*models.Dashboard, error) {
	dashboard := models.Dashboard{
		Name: name,
	}

	if err := d.db.Create(&dashboard).Error; err != nil {
		return nil, err
	}

	return &dashboard, nil
}

func (d *Database) CreateComponent(dashboard *models.Dashboard, name string, cType models.ComponentType, configs interface{}, query string) (*models.Component, error) {
	cfg, err := json.Marshal(configs)
	if err != nil {
		return nil, err
	}

	c := models.Component{
		Dashboard:   *dashboard,
		DashboardID: dashboard.ID,
		Name:        name,
		Query:       query,
		Type:        cType,
		Configs:     cfg,
	}

	if err := d.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (d *Database) UpdateDashboard(dashboard *models.Dashboard) error {
	r := d.db.Save(&dashboard)

	return r.Error
}

func (d *Database) UpdateComponent(component *models.Component) error {
	r := d.db.Save(&component)

	return r.Error
}

func (d *Database) GetComponent(dashboardID, componentID uint) (*models.Component, error) {
	var component models.Component

	if err := d.db.First(&component, componentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if component.DashboardID != dashboardID {
		return nil, nil
	}

	return &component, nil
}

func (d *Database) RunQuery(query string) ([]map[string]interface{}, error) {
	rows, err := d.db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		cm := make([]interface{}, len(cols))
		cmp := make([]interface{}, len(cols))
		for i := range cm {
			cmp[i] = &cm[i]
		}

		if err := rows.Scan(cmp...); err != nil {
			return nil, err
		}

		rm := make(map[string]interface{})
		for i, colName := range cols {
			val := cmp[i].(*interface{})
			rm[colName] = *val
		}

		results = append(results, rm)
	}

	return results, nil
}
