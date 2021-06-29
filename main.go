package main

import (
	"github.com/polyanimal/advertising/internal/server"
	"log"
)

func main() {
	app := server.NewServer()

	if err := app.Run("8080"); err != nil {
		log.Fatal(err)
	}
}
