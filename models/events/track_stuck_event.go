package events

import "github.com/vt-d/audite/models"

type TrackStuckEvent struct {
	Type         EventType `json:"type"`
	GuildID      string    `json:"guildId"`
	models.Track `json:"track"`
	ThresholdMS  int `json:"thresholdMs"`
}
