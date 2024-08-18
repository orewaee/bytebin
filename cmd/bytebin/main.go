package main

import (
	"context"
	"github.com/orewaee/bytebin/internal/app/services"
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/controllers"
	"github.com/orewaee/bytebin/internal/logger"
	"github.com/orewaee/bytebin/internal/meta"
	"github.com/orewaee/bytebin/internal/utils"
	"os"
	"os/signal"
)

func main() {
	log, err := logger.NewZerolog()
	if err != nil {
		panic(err)
	}

	if err := config.Load(); err != nil {
		panic(err)
	}

	if err := utils.CheckDir("bins"); err != nil {
		panic(err)
	}
	binRepo := bin.NewDiskBinRepo()

	if err := utils.CheckDir("metas"); err != nil {
		panic(err)
	}
	metaRepo := meta.NewDiskMetaRepo()

	bytebinApi := services.NewBytebinService(binRepo, metaRepo, log)
	if err := bytebinApi.Load(); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer bytebinApi.Unload()

	addr := config.Get().Addr
	restController := controllers.NewRestController(addr, bytebinApi, log)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := restController.Run(); err != nil {
			log.Fatal().Err(err).Send()
			stop <- os.Interrupt
		}
	}()

	log.Info().Msg("press ctrl+c to exit")

	<-stop
	if err := restController.Shutdown(context.TODO()); err != nil {
		log.Fatal().Err(err).Send()
	}
}
