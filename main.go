package main

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
)

func main() {
	ctx := context.Background()

	config := config.LoadConfig("dev", "auth")
	db := database.Conn(&config.DB)

	defer db.Disconnect(ctx)
	log.Println(db)
}
