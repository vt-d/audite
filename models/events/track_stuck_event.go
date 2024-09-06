package events

import "github.com/vt-d/audite/models"

type TrackStuckEvent struct {
	models.Track `json:"track"`
	ThresholdMS  int `json:"thresholdMs"`
}
