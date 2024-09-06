package events

type WebSocketClosedEvent struct {
	Type     EventType `json:"type"`
	GuildID  string    `json:"guildId"`
	Code     int       `json:"code"`
	Reason   string    `json:"reason"`
	ByRemote bool      `json:"byRemote"`
}
