package strava

import (
	"context"
	"encoding/json"
	"sync"

	pkgerrors "github.com/pkg/errors"
	client "github.com/strava/go.strava"
	"go.uber.org/zap"
)

type Strava struct {
	ctx context.Context
	log *zap.Logger
	mu  sync.Mutex
}

const (
	defaultListLimit = 200
)

func NewStrava(ctx context.Context, log *zap.Logger) *Strava {
	return &Strava{
		ctx: ctx,
		log: log,
	}
}

func (s *Strava) GetAllActivities(user *User) ([]*client.ActivitySummary, error) {
	var allActivities []*client.ActivitySummary
	page := 1

	for {
		activities, err := s.GetPageOfActivities(user, page)
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

func (s *Strava) GetActivityDetail(user *User, activityId int64) (*client.ActivityDetailed, error) {
	service := client.NewActivitiesService(user.st)
	activity, err := service.Get(activityId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return activity, nil
}

func (s *Strava) GetActivityLaps(user *User, activityId int64) ([]*client.LapEffortSummary, error) {
	service := client.NewActivitiesService(user.st)
	laps, err := service.ListLaps(activityId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return laps, nil
}

func (s *Strava) GetPageOfActivities(user *User, page int) ([]*client.ActivitySummary, error) {
	service := client.NewCurrentAthleteService(user.st)

	activities, err := service.
		ListActivities().
		Page(page).
		PerPage(defaultListLimit).
		Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "ListActivities")
	}

	return activities, nil
}

func (s *Strava) GetBeforeOfActivities(user *User, before int64) ([]*client.ActivitySummary, error) {
	service := client.NewCurrentAthleteService(user.st)

	activities, err := service.
		ListActivities().
		Before(int(before)).
		PerPage(defaultListLimit).
		Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "ListActivities")
	}

	return activities, nil
}

func (s *Strava) GetGear(user *User, gearId string) (*client.GearDetailed, error) {
	service := client.NewGearService(user.st)
	gear, err := service.Get(gearId).Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return gear, nil
}

func (s *Strava) GetAthlete(user *User, athleteId int64) (*Athlete, error) {
	service := client.NewAthletesService(user.st)
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

func (s *Strava) ParseError(err error) *client.Error {
	var resp client.Error
	if err := json.Unmarshal([]byte(err.Error()), &resp); err != nil {
		return nil
	}
	return &resp
}
