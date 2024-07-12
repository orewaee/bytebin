package config

import (
	"errors"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Addr     string        `yaml:"addr" env-default:":8080"`
	Lifetime time.Duration `yaml:"lifetime" env-default:"1h"`
}

var config Config

func Get() *Config {
	return &config
}

func Load() error {
	path := getConfigPath()
	if path == "" {
		return errors.New("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("config file does not exist: " + path)
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return errors.New("failed to read config: " + err.Error())
	}

	return nil
}

func getConfigPath() string {
	var result string

	flag.StringVar(&result, "config", "config.yaml", "config file path")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return result
}
