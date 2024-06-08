package models

import (
	"time"

	"gorm.io/datatypes"
)

type Activity struct {
	Id                   int64          `json:"id" gorm:"primaryKey;autoIncrement:false"`
	ExternalId           string         `json:"external_id"`
	UploadId             int64          `json:"upload_id"`
	Name                 string         `json:"name"`
	Distance             float64        `json:"distance"`
	MovingTime           int            `json:"moving_time"`
	ElapsedTime          int            `json:"elapsed_time"`
	TotalElevationGain   float64        `json:"total_elevation_gain"`
	Type                 string         `json:"type"`
	StartDate            time.Time      `json:"start_date"`
	StartDateLocal       time.Time      `json:"start_date_local"`
	TimeZone             string         `json:"time_zone"`
	StartLocation        datatypes.JSON `json:"start_lat_lng"`
	EndLocation          datatypes.JSON `json:"end_lat_lng"`
	City                 string         `json:"location_city"`
	State                string         `json:"location_state"`
	Country              string         `json:"location_country"`
	AchievementCount     int            `json:"achievement_count"`
	KudosCount           int            `json:"kudos_count"`
	CommentCount         int            `json:"comment_count"`
	AthleteCount         int            `json:"athlete_count"`
	PhotoCount           int            `json:"photo_count"`
	Map                  datatypes.JSON `json:"map"`
	Trainer              bool           `json:"trainer"`
	Commute              bool           `json:"commute"`
	Manual               bool           `json:"manual"`
	Private              bool           `json:"private"`
	Flagged              bool           `json:"flagged"`
	AverageSpeed         float64        `json:"average_speed"`
	MaximumSpeed         float64        `json:"max_speed"`
	AverageCadence       float64        `json:"average_cadence"`
	AverageTemperature   float64        `json:"average_temp"`
	AveragePower         float64        `json:"average_watts"`
	WeightedAveragePower int            `json:"weighted_average_watts"`
	Kilojoules           float64        `json:"kilojoules"`
	DeviceWatts          bool           `json:"device_watts"`
	AverageHeartRate     float64        `json:"average_heart_rate"`
	MaximumHeartRate     float64        `json:"max_heart_rate"`
	Truncated            int            `json:"truncated"` // only present if activity is owned by authenticated athlete, returns 0 if not truncated by privacy zones
	HasKudos             bool           `json:"has_kudos"`
	AthleteID            uint           `json:"athlete_id"`
	Athlete              Athlete        `gorm:"foreignkey:AthleteID" json:"athlete"`
	GearID               *string        `json:"gear_id"`
	Gear                 *Gear          `gorm:"foreignkey:GearID" json:"gear,omitempty"`
}
