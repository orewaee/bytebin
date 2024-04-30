package main

import (
	"github.com/orewaee/bytebin/internal/app"
	"github.com/orewaee/bytebin/internal/config"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[bytebin] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	bytebin := app.New(
		&config.Config{
			Addr:     ":8080",
			Limit:    1024,
			Lifetime: time.Second * 10,
		},
		logger,
	)

	if err := bytebin.Run(); err != nil {
		logger.Fatalln(err)
	}
}
