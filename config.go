package main

// unfinished

type ServerConfig struct {
	Port    int16  `toml:"port"`
	Address string `toml:"address"`
}

type Options struct {
	StatsReload int `toml:"stats_reload,omitempty"`
}

type Config struct {
	ServerConfig `toml:"server"`
	Options      `toml:"opts"`
}
