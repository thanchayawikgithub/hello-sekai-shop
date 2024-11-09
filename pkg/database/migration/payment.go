package migration

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func paymentDBConn(cfg *config.Config) *mongo.Database {
	return database.Conn(&cfg.DB).Database("payment_db")
}

func PaymentMigrate(ctx context.Context, cfg *config.Config) {
	db := paymentDBConn(cfg)
	defer db.Client().Disconnect(ctx)

	//payments
	col := db.Collection("payment_queue")

	results, err := col.InsertOne(ctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate payment completed: ", results)
}
