package events

import "github.com/vt-d/audite/models"

type TrackStartEvent struct {
	Track models.Track `json:"track"`
}
