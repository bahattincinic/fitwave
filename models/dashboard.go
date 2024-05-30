package models

import (
	"time"

	"gorm.io/datatypes"
)

type Dashboard struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`
}

type ComponentType string

const (
	TableComponent     = ComponentType("table")
	PieChartComponent  = ComponentType("pie_chart")
	BarChartComponent  = ComponentType("bar_chart")
	LineChartComponent = ComponentType("line_chart")
)

type Component struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DashboardID uint
	Dashboard   Dashboard `gorm:"foreignkey:DashboardID" json:"dashboard"`

	Name    string         `json:"name"`
	Query   string         `json:"query"`
	Type    ComponentType  `json:"type"`
	Configs datatypes.JSON `json:"configs"`
}
