package strava

import client "github.com/strava/go.strava"

type Athlete struct {
	Athlete *client.AthleteSummary
	Stats   *client.AthleteStats
}
