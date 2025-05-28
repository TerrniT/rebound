package main

import (
	"log"

	"github.com/terrnit/rebound/backend/config"
	_ "github.com/terrnit/rebound/backend/docs" // This will be generated
	app "github.com/terrnit/rebound/backend/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
