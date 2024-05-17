package importer

import (
	"context"
	"encoding/json"

	"github.com/bahattincinic/fitwave/config"
	"github.com/bahattincinic/fitwave/database"
	"github.com/bahattincinic/fitwave/models"
	"github.com/bahattincinic/fitwave/strava"
	pkgerrors "github.com/pkg/errors"
	client "github.com/strava/go.strava"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Importer struct {
	ctx context.Context
	cfg *config.Config
	st  *strava.Strava
	log *zap.Logger
	db  *database.Database
}

func NewImporter(ctx context.Context, cfg *config.Config, log *zap.Logger, st *strava.Strava, db *database.Database) *Importer {
	return &Importer{
		ctx: ctx,
		cfg: cfg,
		log: log,
		st:  st,
		db:  db,
	}
}

func (im *Importer) updateGears(tx *gorm.DB, gears []client.GearDetailed, gearAthletes map[string]int64) error {
	var rows []models.Gear

	for _, gear := range gears {
		rows = append(rows, models.Gear{
			Id:          gear.Id,
			Name:        gear.Name,
			Primary:     gear.Primary,
			Distance:    gear.Distance,
			BrandName:   gear.BrandName,
			ModelName:   gear.ModelName,
			FrameType:   gear.FrameType.String(),
			Description: gear.Description,
			AthleteID:   gearAthletes[gear.Id],
		})
	}
	return im.db.UpsertGears(tx, rows)
}

func (im *Importer) updateAthletes(tx *gorm.DB, athletes []strava.Athlete) error {
	var rows []models.Athlete

	for _, athlete := range athletes {
		sum := athlete.Athlete
		stats, err := json.Marshal(athlete.Stats)
		if err != nil {
			return pkgerrors.Wrap(err, "Marshal")
		}

		rows = append(rows, models.Athlete{
			Id:               sum.Id,
			FirstName:        sum.FirstName,
			LastName:         sum.LastName,
			ProfileMedium:    sum.ProfileMedium,
			Profile:          sum.Profile,
			City:             sum.City,
			State:            sum.State,
			Country:          sum.Country,
			Gender:           string(sum.Gender),
			Friend:           sum.Friend,
			Follower:         sum.Follower,
			Premium:          sum.Premium,
			CreatedAt:        sum.CreatedAt,
			UpdatedAt:        sum.UpdatedAt,
			ApproveFollowers: sum.ApproveFollowers,
			BadgeTypeId:      sum.BadgeTypeId,
			Stats:            stats,
		})
	}

	return im.db.UpsertAthletes(tx, rows)
}

func (im *Importer) updateActivities(tx *gorm.DB, activities []*client.ActivitySummary) error {
	var rows []models.Activity

	for _, activity := range activities {
		actMap, err := json.Marshal(activity.Map)
		if err != nil {
			return pkgerrors.Wrap(err, "Marshal")
		}

		rows = append(rows, models.Activity{
			Id:                   activity.Id,
			ExternalId:           activity.ExternalId,
			UploadId:             activity.UploadId,
			AthleteID:            uint(activity.Athlete.Id),
			Name:                 activity.Name,
			Distance:             activity.Distance,
			MovingTime:           activity.MovingTime,
			ElapsedTime:          activity.ElapsedTime,
			TotalElevationGain:   activity.TotalElevationGain,
			Type:                 string(activity.Type),
			StartDate:            activity.StartDate,
			StartDateLocal:       activity.StartDateLocal,
			TimeZone:             activity.TimeZone,
			StartLocation:        activity.StartLocation.String(),
			EndLocation:          activity.EndLocation.String(),
			City:                 activity.City,
			State:                activity.State,
			Country:              activity.Country,
			AchievementCount:     activity.AchievementCount,
			KudosCount:           activity.KudosCount,
			CommentCount:         activity.CommentCount,
			AthleteCount:         activity.AthleteCount,
			PhotoCount:           activity.PhotoCount,
			Map:                  actMap,
			Trainer:              activity.Trainer,
			Commute:              activity.Commute,
			Manual:               activity.Manual,
			Private:              activity.Private,
			Flagged:              activity.Flagged,
			GearId:               activity.GearId,
			AverageSpeed:         activity.AverageSpeed,
			MaximumSpeed:         activity.MaximunSpeed,
			AverageCadence:       activity.AverageCadence,
			AverageTemperature:   activity.AverageTemperature,
			AveragePower:         activity.AveragePower,
			WeightedAveragePower: activity.WeightedAveragePower,
			Kilojoules:           activity.Kilojoules,
			DeviceWatts:          activity.DeviceWatts,
			AverageHeartRate:     activity.AverageHeartrate,
			MaximumHeartRate:     activity.MaximumHeartrate,
			Truncated:            activity.Truncated,
			HasKudos:             activity.HasKudoed,
		})
	}

	return im.db.UpsertActivities(tx, rows)
}

func (im *Importer) Import() error {
	activities, err := im.st.GetAllActivities()
	if err != nil {
		return pkgerrors.Wrap(err, "GetAllActivities")
	}

	im.log.Info("activities have been fetched",
		zap.Int("count", len(activities)))

	athleteIds := make(map[int64]bool)
	gearIds := make(map[string]bool)
	gearAthletes := make(map[string]int64)

	var athletes []strava.Athlete
	var gears []client.GearDetailed

	for _, act := range activities {
		athleteId := act.Athlete.Id
		gearId := act.GearId

		if _, ok := athleteIds[athleteId]; !ok {
			athleteIds[athleteId] = true
			athlete, err := im.st.GetAthlete(athleteId)
			if err != nil {
				im.log.Info("could not fetch athlete",
					zap.Int64("id", athleteId),
					zap.Error(err))

				return pkgerrors.Wrap(err, "GetAthlete")
			}
			athletes = append(athletes, *athlete)
		}

		if _, ok := gearIds[gearId]; !ok && gearId != "" {
			gearIds[gearId] = true
			gear, err := im.st.GetGear(gearId)
			if err != nil {
				im.log.Info("could not fetch gear",
					zap.String("id", gearId),
					zap.Error(err))

				return pkgerrors.Wrap(err, "GetGear")
			}
			gearAthletes[gearId] = athleteId
			gears = append(gears, *gear)
		}
	}

	im.log.Info("athletes have been fetched",
		zap.Int("count", len(athletes)))

	im.log.Info("gears have been fetched",
		zap.Int("count", len(gears)))

	tx := im.db.BeginTransaction()
	defer tx.Rollback()

	if len(gears) > 0 {
		if err := im.updateGears(tx, gears, gearAthletes); err != nil {
			return pkgerrors.Wrap(err, "updateGears")
		}
	}

	if len(athletes) > 0 {
		if err := im.updateAthletes(tx, athletes); err != nil {
			return pkgerrors.Wrap(err, "updateAthletes")
		}
	}

	if len(activities) > 0 {
		if err := im.updateActivities(tx, activities); err != nil {
			return pkgerrors.Wrap(err, "updateActivities")
		}
	}
	return tx.Commit().Error
}
