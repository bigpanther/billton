package main

import (
	"log"

	"github.com/bigpanther/billton/internal/app"
)

func main() {
	server, err := app.App()
	if err != nil {
		log.Fatal("failed to initialize app", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
