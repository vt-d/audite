package events

import "github.com/vt-d/audite/models"

type Severity string

const (
	SeverityCommon     = "common"
	SeveritySuspicious = "suspicious"
	SeverityFault      = "fault"
)

type Exception struct {
	Severity
	Message string
	Cause   string
}

type TrackExceptionEvent struct {
	Type         EventType `json:"type"`
	GuildID      string    `json:"guildId"`
	models.Track `json:"track"`
	Exception    `json:"exception"`
}
