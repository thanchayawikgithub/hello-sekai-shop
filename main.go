package main

import (
	"context"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"github.com/thanchayawikgithub/hello-sekai-shop/server"
)

func main() {
	ctx := context.Background()

	config := config.LoadConfig("dev", "auth")

	db := database.Conn(&config.DB)
	defer db.Disconnect(ctx)

	server.Start(ctx, config, db)
}
