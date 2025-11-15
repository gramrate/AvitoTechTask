package main

import (
	"AvitoTechTask/internal/adapters/app"
	"AvitoTechTask/internal/adapters/controller/api/server"
	"log"
)

func main() {
	mainApp, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	server.Setup(mainApp)
	mainApp.Start()
}
