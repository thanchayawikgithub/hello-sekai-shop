package main

import (
	"context"
	"log"
	"os"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database/migration"
)

func main() {
	ctx := context.Background()

	// Check if the state and service arguments are provided
	if len(os.Args) < 3 {
		log.Fatal("Error: state and service arguments are required")
	}

	state := os.Args[1]   // e.g., "dev"
	service := os.Args[2] // e.g., "auth"

	// Load configuration based on state and service
	cfg := config.LoadConfig(state, service)

	switch cfg.App.Name {
	case "auth":
		migration.AuthMigrate(ctx, cfg)
	case "inventory":
	case "item":
	case "payment":
	case "player":
	}
}
