package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/orewaee/bytebin/internal/app"
	"github.com/orewaee/bytebin/internal/config"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[bytebin] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		logger.Fatalln(err)
	}

	bytebin := app.New(&cfg, logger)

	if err := bytebin.Run(); err != nil {
		logger.Fatalln(err)
	}
}
