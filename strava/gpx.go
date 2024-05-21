package strava

import (
	"encoding/xml"
	"time"

	pkgerrors "github.com/pkg/errors"
	client "github.com/strava/go.strava"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	XMLNS   string   `xml:"xmlns,attr"`
	Trk     Trk      `xml:"trk"`
}

type Trk struct {
	Name   string  `xml:"name"`
	Trkseg []Trkpt `xml:"trkseg>trkpt"`
}

type Trkpt struct {
	Lat   float64 `xml:"lat,attr"`
	Lon   float64 `xml:"lon,attr"`
	Ele   float64 `xml:"ele,omitempty"`
	Time  string  `xml:"time"`
	Temp  int     `xml:"temp,omitempty"`
	HR    int     `xml:"hr,omitempty"`
	Cad   int     `xml:"cad,omitempty"`
	Pwr   int     `xml:"pwr,omitempty"`
	Spd   float64 `xml:"spd,omitempty"`
	Dist  float64 `xml:"dist,omitempty"`
	Mov   bool    `xml:"mov,omitempty"`
	Grade float64 `xml:"grade,omitempty"`
}

func (s *Strava) GetActivityStream(user *User, activityId int64) (*client.StreamSet, error) {
	types := []client.StreamType{
		client.StreamTypes.Time,
		client.StreamTypes.Location,
		client.StreamTypes.Distance,
		client.StreamTypes.Elevation,
		client.StreamTypes.Speed,
		client.StreamTypes.HeartRate,
		client.StreamTypes.Cadence,
		client.StreamTypes.Power,
		client.StreamTypes.Temperature,
		client.StreamTypes.Moving,
		client.StreamTypes.Grade,
	}

	service := client.NewActivityStreamsService(user.st)
	streams, err := service.Get(activityId, types).Do()

	if err != nil {
		parsedErr := s.ParseError(err)
		if parsedErr != nil && parsedErr.Message == "Resource Not Found" {
			return nil, nil
		}
		return nil, pkgerrors.Wrap(err, "Get")
	}

	return streams, nil
}
func (s *Strava) ExportGPX(user *User, activityId int64) (string, error) {
	streams, err := s.GetActivityStream(user, activityId)
	if err != nil {
		return "", err
	}

	if streams == nil {
		return "", nil
	}

	var trkpts []Trkpt
	for i := 0; i < len(streams.Time.Data); i++ {
		trkpt := Trkpt{
			Time: time.Unix(int64(streams.Time.Data[i]), 0).UTC().Format(time.RFC3339),
		}

		if streams.Location != nil && len(streams.Location.Data) > i {
			trkpt.Lat = streams.Location.Data[i][0]
			trkpt.Lon = streams.Location.Data[i][1]
		}

		if streams.Elevation != nil && len(streams.Elevation.Data) > i {
			trkpt.Ele = streams.Elevation.Data[i]
		}

		if streams.Temperature != nil && len(streams.Temperature.Data) > i {
			trkpt.Temp = streams.Temperature.Data[i]
		}

		if streams.HeartRate != nil && len(streams.HeartRate.Data) > i {
			trkpt.HR = streams.HeartRate.Data[i]
		}

		if streams.Cadence != nil && len(streams.Cadence.Data) > i {
			trkpt.Cad = streams.Cadence.Data[i]
		}

		if streams.Power != nil && len(streams.Power.Data) > i {
			trkpt.Pwr = streams.Power.Data[i]
		}

		if streams.Speed != nil && len(streams.Speed.Data) > i {
			trkpt.Spd = streams.Speed.Data[i]
		}

		if streams.Distance != nil && len(streams.Distance.Data) > i {
			trkpt.Dist = streams.Distance.Data[i]
		}

		if streams.Moving != nil && len(streams.Moving.Data) > i {
			trkpt.Mov = streams.Moving.Data[i]
		}

		if streams.Grade != nil && len(streams.Grade.Data) > i {
			trkpt.Grade = streams.Grade.Data[i]
		}

		trkpts = append(trkpts, trkpt)
	}

	gpx := GPX{
		Version: "1.1",
		Creator: "Strava GPX Exporter",
		XMLNS:   "http://www.topografix.com/GPX/1/1",
		Trk: Trk{
			Name:   "Strava Activity",
			Trkseg: trkpts,
		},
	}

	output, err := xml.MarshalIndent(gpx, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
