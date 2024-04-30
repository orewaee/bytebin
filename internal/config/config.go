package config

import (
	"time"
)

type Config struct {
	Addr     string        `env:"ADDR" env-default:":8080"`
	Limit    int64         `env:"LIMIT" env-default:"104857600"`
	Lifetime time.Duration `env:"LIFETIME" env-default:"1h"`
}
