package models

type TrackInfo struct {
	Identifier string `json:"identifier"`
	Seekable   bool   `json:"seekable"`
	Author     string `json:"author"`
	Length     int    `json:"length"`
	IsStream   bool   `json:"isStream"`
	Position   int    `json:"position"`
	Title      string `json:"title"`
	URI        string `json:"uri,omitempty"`

	ArtworkURL string `json:"artworkUrl,omitempty"`
	Isrc       string `json:"isrc,omitempty"`
	SourceName string `json:"sourceName"`
}

type Track struct {
	Encoded    string                 `json:"encoded"`
	Info       TrackInfo              `json:"info"`
	PluginInfo map[string]interface{} `json:"pluginInfo"`
	UserData   map[string]interface{} `json:"userData"`
}
