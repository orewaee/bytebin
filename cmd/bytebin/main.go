package main

import (
	"flag"
	"github.com/orewaee/bytebin/internal/app"
	"log"
	"os"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	logger := log.New(os.Stdout, "[bytebin] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	bytebin := app.New(":"+strconv.Itoa(*port), logger)

	if err := bytebin.Run(); err != nil {
		logger.Fatalln(err)
	}
}
