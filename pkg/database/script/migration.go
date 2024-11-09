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

	if len(os.Args) < 3 {
		log.Fatal("Error: state and service arguments are required")
	}

	state := os.Args[1]
	service := os.Args[2]

	cfg := config.LoadConfig(state, service)

	switch cfg.App.Name {
	case "auth":
		migration.AuthMigrate(ctx, cfg)
	case "inventory":
		migration.InventoryMigrate(ctx, cfg)
	case "item":
		migration.ItemMigrate(ctx, cfg)
	case "payment":
		migration.PaymentMigrate(ctx, cfg)
	case "player":
		migration.PlayerMigrate(ctx, cfg)
	}
}
