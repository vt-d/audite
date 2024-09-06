package events

import "github.com/vt-d/audite/models"

type TrackEndReason string

const (
	TrackEndReasonFinished   TrackEndReason = "finished"
	TrackEndReasonLoadFailed TrackEndReason = "loadFailed"
	TrackEndReasonStopped    TrackEndReason = "stopped"
	TrackEndReasonReplaced   TrackEndReason = "replaced"
	TrackEndReasonCleanup    TrackEndReason = "cleanup"
)

type TrackEndEvent struct {
	Type         EventType `json:"type"`
	GuildID      string    `json:"guildId"`
	models.Track `json:"track"`
	Reason       TrackEndReason `json:"reason"`
}
