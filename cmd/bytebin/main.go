package main

import (
	"github.com/orewaee/bytebin/internal/app"
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/config"
	"github.com/orewaee/bytebin/internal/logger"
	"github.com/orewaee/bytebin/internal/meta"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/orewaee/bytebin/internal/utils"
)

func main() {
	log, err := logger.New(".")
	if err != nil {
		panic(err)
	}

	if err := config.Load(); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := utils.CheckDir("metas"); err != nil {
		log.Fatal().Err(err).Send()
	}
	diskMetas := meta.NewDiskManager()

	if err := utils.CheckDir("bins"); err != nil {
		log.Fatal().Err(err).Send()
	}
	diskBins := bin.NewDiskManager()

	diskStorage := storage.NewDiskStorage(diskBins, diskMetas)
	if err := diskStorage.Load(); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() {
		if err := diskStorage.Unload(); err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	bytebin := app.New(diskStorage, log)
	if err := bytebin.Run(config.Get().Addr); err != nil {
		log.Fatal().Err(err).Send()
	}
}
