package main

import (
	"github.com/orewaee/bytebin/internal/app"
	"github.com/orewaee/bytebin/internal/bin"
	"github.com/orewaee/bytebin/internal/meta"
	"github.com/orewaee/bytebin/internal/storage"
	"github.com/orewaee/bytebin/internal/utils"
	"log"
)

func main() {
	if err := utils.CheckDir("metas"); err != nil {
		log.Fatalln(err)
	}
	diskMetas := meta.NewDiskManager()

	if err := utils.CheckDir("bins"); err != nil {
		log.Fatalln(err)
	}
	diskBins := bin.NewDiskManager()

	diskStorage := storage.NewDiskStorage(diskBins, diskMetas)
	if err := diskStorage.Load(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := diskStorage.Unload(); err != nil {
			log.Fatalln(err)
		}
	}()

	bytebin := app.New(diskStorage)
	if err := bytebin.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
