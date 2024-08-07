package main

import (
	"github.com/orewaee/bytebin/internal/adapters/bin"
	"github.com/orewaee/bytebin/internal/adapters/http"
	"github.com/orewaee/bytebin/internal/adapters/meta"
	"github.com/orewaee/bytebin/internal/app/services"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/logger"
	"github.com/orewaee/bytebin/internal/utils"
	"os"
	"os/signal"
)

func main() {
	log, err := logger.New(".")
	if err != nil {
		panic(err)
	}

	if err := config.Load(); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := utils.CheckDir("bins"); err != nil {
		log.Fatal().Err(err).Send()
	}
	binRepo := bin.NewDiskBinRepo()

	if err := utils.CheckDir("metas"); err != nil {
		log.Fatal().Err(err).Send()
	}
	metaRepo := meta.NewDiskMetaRepo()

	bytebin := services.NewBytebinService(binRepo, metaRepo, log)
	if err := bytebin.Load(); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer bytebin.Unload()

	server := http.NewServer(bytebin, log)
	addr := config.Get().Addr

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := server.Run(addr); err != nil {
			log.Fatal().Err(err).Send()
			stop <- os.Interrupt
		}
	}()

	log.Info().Msg("Press Ctrl+C to exit")

	<-stop
	if err := server.Shutdown(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
