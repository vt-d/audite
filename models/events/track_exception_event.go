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
	models.Track `json:"track"`
	Exception    `json:"exception"`
}
