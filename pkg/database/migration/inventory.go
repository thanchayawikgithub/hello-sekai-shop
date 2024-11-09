package migration

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func inventoryDBConn(cfg *config.Config) *mongo.Database {
	return database.Conn(&cfg.DB).Database("inventory_db")
}

func InventoryMigrate(ctx context.Context, cfg *config.Config) {
	db := inventoryDBConn(cfg)
	defer db.Client().Disconnect(ctx)

	//inventories
	col := db.Collection(database.PlayerInventoryCollection)
	indexes, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "player_id", Value: 1}, {Key: "item_id", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	col = db.Collection(database.PlayerInventoryQueueCollection)
	results, err := col.InsertOne(ctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate inventory completed: ", results)
}
