package config

import (
	"errors"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Addr     string        `yaml:"addr" env:"ADDR" env-default:":8080"`
	Lifetime time.Duration `yaml:"lifetime" env:"LIFETIME" env-default:"1h"`
}

var config Config

func Get() *Config {
	return &config
}

func Load() error {
	path := getConfigPath()

	if path == "" {
		if err := cleanenv.ReadEnv(&config); err != nil {
			return err
		}

		return nil
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return errors.New("failed to read config: " + err.Error())
	}

	return nil
}

func getConfigPath() string {
	var result string

	flag.StringVar(&result, "config", "", "config file path")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return result
}
