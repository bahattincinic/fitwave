package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// createStravaTables returns the migrations for creating Strava-related tables.
func (m *Migration) createStravaTables() []*gormigrate.Migration {
	return []*gormigrate.Migration{{
		ID: "202404171309",
		Migrate: func(tx *gorm.DB) error {
			type athlete struct {
				Id               int64 `gorm:"primaryKey;autoIncrement:false"`
				FirstName        string
				LastName         string
				ProfileMedium    string
				Profile          string
				City             string
				State            string
				Country          string
				Gender           string
				Friend           string
				Follower         string
				Premium          bool
				CreatedAt        time.Time
				UpdatedAt        time.Time
				ApproveFollowers bool
				BadgeTypeId      int
				Stats            datatypes.JSON
			}
			type gear struct {
				Id          string `gorm:"primaryKey;autoIncrement:false"`
				Name        string
				Primary     bool
				Distance    float64
				BrandName   string
				ModelName   string
				Type        string
				Description string
				AthleteID   int64
				Athlete     *athlete `gorm:"foreignkey:AthleteID"`
			}
			type activity struct {
				Id                   int64 `gorm:"primaryKey;autoIncrement:false"`
				ExternalId           string
				UploadId             int64
				Name                 string
				Distance             float64
				MovingTime           int
				ElapsedTime          int
				TotalElevationGain   float64
				Type                 string
				StartDate            time.Time
				StartDateLocal       time.Time
				TimeZone             string
				StartLocation        datatypes.JSON
				EndLocation          datatypes.JSON
				City                 string
				State                string
				Country              string
				AchievementCount     int
				KudosCount           int
				CommentCount         int
				AthleteCount         int
				PhotoCount           int
				Map                  datatypes.JSON
				Trainer              bool
				Commute              bool
				Manual               bool
				Private              bool
				Flagged              bool
				AverageSpeed         float64
				MaximumSpeed         float64
				AverageCadence       float64
				AverageTemperature   float64
				AveragePower         float64
				WeightedAveragePower int
				Kilojoules           float64
				DeviceWatts          bool
				AverageHeartRate     float64
				MaximumHeartRate     float64
				Truncated            int
				HasKudos             bool
				AthleteID            uint
				Athlete              athlete `gorm:"foreignkey:AthleteID"`
				GearID               *string
				Gear                 *athlete `gorm:"foreignkey:GearID"`
			}
			return tx.AutoMigrate(&activity{}, &athlete{}, &gear{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("activities", "gears", "athletes")
		},
	}}
}
