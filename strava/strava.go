package strava

import (
	"context"

	"github.com/bahattincinic/fitwave/config"
	pkgerrors "github.com/pkg/errors"
	client "github.com/strava/go.strava"
	"go.uber.org/zap"
)

type Strava struct {
	ctx context.Context
	cfg *config.Config
	log *zap.Logger
	st  *client.Client
}

const (
	defaultListLimit = 200
)

func NewStrava(ctx context.Context, cfg *config.Config, log *zap.Logger) *Strava {
	st := client.NewClient(cfg.Strava.AccessToken)

	return &Strava{
		ctx: ctx,
		cfg: cfg,
		log: log,
		st:  st,
	}
}

func (s *Strava) GetAllActivities() ([]*client.ActivitySummary, error) {
	var allActivities []*client.ActivitySummary
	page := 1

	for {
		activities, err := s.GetPageOfActivities(page)
		if err != nil {
			return nil, err
		}

		if len(activities) == 0 {
			// No more activities to fetch
			break
		}

		allActivities = append(allActivities, activities...)
		page++
	}

	return allActivities, nil
}

func (s *Strava) GetActivityDetail(activityId int64) (*client.ActivityDetailed, error) {
	service := client.NewActivitiesService(s.st)
	activity, err := service.Get(activityId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return activity, nil
}

func (s *Strava) GetPageOfActivities(page int) ([]*client.ActivitySummary, error) {
	service := client.NewAthletesService(s.st)

	activities, err := service.ListActivities(s.cfg.Strava.AthleteId).
		Page(page).
		PerPage(defaultListLimit).
		Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "ListActivities")
	}

	return activities, nil
}

func (s *Strava) GetBeforeOfActivities(before int64) ([]*client.ActivitySummary, error) {
	service := client.NewAthletesService(s.st)

	activities, err := service.ListActivities(s.cfg.Strava.AthleteId).
		Before(before).
		PerPage(defaultListLimit).
		Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "ListActivities")
	}

	return activities, nil
}

func (s *Strava) GetGear(gearId string) (*client.GearDetailed, error) {
	service := client.NewGearService(s.st)
	gear, err := service.Get(gearId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return gear, nil
}

func (s *Strava) GetAthlete(athleteId int64) (*Athlete, error) {
	service := client.NewAthletesService(s.st)
	athlete, err := service.Get(athleteId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	stats, err := service.Stats(athleteId).Do()
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Stats.Get")
	}

	return &Athlete{
		Athlete: athlete,
		Stats:   stats,
	}, nil
}
