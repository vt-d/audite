package models

type Ready struct {
	Op        string `json:"op"`
	SessionID string `json:"session_id"`
	Resumed   bool   `json:"resumed"`
}

type PlayerState struct {
	Time      int64 `json:"time"`
	Position  int64 `json:"position"`
	Connected bool  `json:"connected"`
	Ping      int   `json:"ping"`
}

type PlayerUpdate struct {
	GuildID int         `json:"guildId"`
	State   PlayerState `json:"state"`
}

type FrameStats struct {
	Sent    int `json:"sent"`
	Nulled  int `json:"nulled"`
	Deficit int `json:"deficit"`
}

type CPUStats struct {
	Cores      int     `json:"cores"`
	SystemLoad float64 `json:"systemLoad"`
	AppLoad    float64 `json:"lavalinkLoad"`
}

type MemoryStats struct {
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
	Allocated  uint64 `json:"allocated"`
	Reservable uint64 `json:"reservable"`
}

type Stats struct {
	Players        int `json:"players"`
	PlayingPlayers int `json:"playingPlayers"`
	Uptime         int `json:"uptime"`
	MemoryStats    `json:"memory"`
	CPUStats       `json:"cpu"`
	FrameStats     `json:"frameStats"`
}
