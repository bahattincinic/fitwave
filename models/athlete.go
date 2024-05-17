package models

import (
	"time"

	"gorm.io/datatypes"
)

type Athlete struct {
	Id               int64          `json:"id" gorm:"primaryKey;autoIncrement:false"`
	FirstName        string         `json:"firstname"`
	LastName         string         `json:"lastname"`
	ProfileMedium    string         `json:"profile_medium"` // URL to a 62x62 pixel profile picture
	Profile          string         `json:"profile"`        // URL to a 124x124 pixel profile picture
	City             string         `json:"city"`
	State            string         `json:"state"`
	Country          string         `json:"country"`
	Gender           string         `json:"sex"`
	Friend           string         `json:"friend"`   // ‘pending’, ‘accepted’, ‘blocked’ or ‘null’, the authenticated athlete’s following status of this athlete
	Follower         string         `json:"follower"` // this athlete’s following status of the authenticated athlete
	Premium          bool           `json:"premium"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	ApproveFollowers bool           `json:"approve_followers"` // if has enhanced privacy enabled
	BadgeTypeId      int            `json:"badge_type_id"`
	Stats            datatypes.JSON `json:"tx"`
}
