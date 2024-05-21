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

func (im *Importer) updateGears(tx *gorm.DB, gears []client.GearDetailed, user *strava.User) error {
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
			AthleteID:   user.Athlete.Id,
		})
	}
	return im.db.UpsertGears(tx, rows)
}

func (im *Importer) updateAthlete(tx *gorm.DB, athlete strava.Athlete) error {
	sum := athlete.Athlete
	stats, err := json.Marshal(athlete.Stats)
	if err != nil {
		return pkgerrors.Wrap(err, "Marshal")
	}

	row := models.Athlete{
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
	}

	return im.db.UpsertAthlete(tx, &row)
}

func (im *Importer) updateActivities(tx *gorm.DB, activities []*client.ActivitySummary, user *strava.User) error {
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
			GearID:               activity.GearId,
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

func (im *Importer) Import(user *strava.User) error {
	activities, err := im.st.GetAllActivities(user)
	if err != nil {
		return pkgerrors.Wrap(err, "GetAllActivities")
	}

	im.log.Info("activities have been fetched",
		zap.Int("count", len(activities)))

	gearIds := make(map[string]bool)

	var athletes []strava.Athlete
	var gears []client.GearDetailed

	athlete, err := im.st.GetAthlete(user, user.Athlete.Id)
	if err != nil {
		im.log.Info("could not fetch athlete",
			zap.Int64("id", user.Athlete.Id),
			zap.Error(err))
		return pkgerrors.Wrap(err, "GetAthlete")
	}

	for _, act := range activities {
		gearId := act.GearId

		if _, ok := gearIds[gearId]; !ok && gearId != "" {
			gearIds[gearId] = true
			gear, err := im.st.GetGear(user, gearId)
			if err != nil {
				im.log.Info("could not fetch gear",
					zap.String("id", gearId),
					zap.Error(err))

				return pkgerrors.Wrap(err, "GetGear")
			}
			gears = append(gears, *gear)
		}
	}

	im.log.Info("athletes have been fetched",
		zap.Int("count", len(athletes)))

	im.log.Info("gears have been fetched",
		zap.Int("count", len(gears)))

	tx := im.db.BeginTransaction()
	defer tx.Rollback()

	if err := im.updateAthlete(tx, *athlete); err != nil {
		return pkgerrors.Wrap(err, "updateAthlete")
	}

	if len(gears) > 0 {
		if err := im.updateGears(tx, gears, user); err != nil {
			return pkgerrors.Wrap(err, "updateGears")
		}
	}

	if len(activities) > 0 {
		if err := im.updateActivities(tx, activities, user); err != nil {
			return pkgerrors.Wrap(err, "updateActivities")
		}
	}
	return tx.Commit().Error
}
