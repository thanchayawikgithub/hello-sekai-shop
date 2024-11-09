package main

import (
	"context"
	"os"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"github.com/thanchayawikgithub/hello-sekai-shop/server"
)

func main() {
	ctx := context.Background()

	// Get environment variables
	state := os.Getenv("STATE")
	service := os.Getenv("SERVICE")

	// Load configuration based on state and service
	cfg := config.LoadConfig(state, service)

	// Initialize database connection and server with the loaded configuration
	db := database.Conn(&cfg.DB)
	defer db.Disconnect(ctx)

	// Start the server with context, config, and database connection
	server.Start(ctx, cfg, db)
}
