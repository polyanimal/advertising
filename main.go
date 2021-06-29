package main

import (
	"github.com/polyanimal/advertising/internal/server"
	constants "github.com/polyanimal/advertising/pkg/const"
	"log"
)

func main() {
	app := server.NewServer()

	if err := app.Run(constants.Port); err != nil {
		log.Fatal(err)
	}
}
