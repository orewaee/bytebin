package main

import (
	"flag"
	"github.com/orewaee/bytebin/internal/app"
	"log"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	err := app.New(":" + strconv.Itoa(*port)).Run()
	if err != nil {
		log.Fatalln(err)
	}
}
