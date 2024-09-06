package events

type EventType string

const (
	OpTrackStartEvent      EventType = "TrackStartEvent"
	OpTrackEndEvent        EventType = "TrackEndEvent"
	OpTrackExceptionEvent  EventType = "TrackExceptionEvent"
	OpTrackStuckEvent      EventType = "TrackStuckEvent"
	OpWebSocketClosedEvent EventType = "WebSocketClosedEvent"
)

type BaseEvent struct {
	Type    EventType `json:"type"`
	GuildID string    `json:"guildId"`
}
