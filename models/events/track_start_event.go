package events

import "github.com/vt-d/audite/models"

type TrackStartEvent struct {
	Type    EventType    `json:"type"`
	GuildID string       `json:"guildId"`
	Track   models.Track `json:"track"`
}
