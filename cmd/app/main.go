package main

import (
	"log"

	config "github.com/audryus/boleto-palm-tree/configs"
	"github.com/audryus/boleto-palm-tree/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
