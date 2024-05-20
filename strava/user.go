package strava

import (
	"github.com/bahattincinic/fitwave/models"
	pkgerrors "github.com/pkg/errors"
	client "github.com/strava/go.strava"
)

type User struct {
	st          *client.Client
	AccessToken string
	cfg         *models.Config
	Athlete     *client.AthleteDetailed
}

func (s *Strava) NewUser(cfg *models.Config, accessToken string) (*User, error) {
	athlete, err := s.GetCurrentAthlete(accessToken)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "GetCurrentAthlete")
	}

	return &User{
		st:          client.NewClient(accessToken),
		Athlete:     athlete,
		AccessToken: accessToken,
		cfg:         cfg,
	}, nil
}

func (s *Strava) GetCurrentAthlete(accessToken string) (*client.AthleteDetailed, error) {
	cl := client.NewClient(accessToken)
	service := client.NewCurrentAthleteService(cl)
	athlete, err := service.Get().Do()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return athlete, nil
}
