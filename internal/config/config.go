package config

import "time"

type Config struct {
	Addr     string
	Limit    int
	Lifetime time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Addr:     ":8080",
		Limit:    1024,
		Lifetime: 3600,
	}
}
