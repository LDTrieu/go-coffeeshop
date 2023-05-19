package main

import (
	"log"

	"go-coffeeshop/cmd/product/config"
	"go-coffeeshop/internal/product/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
