package migration

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func itemDBConn(cfg *config.Config) *mongo.Database {
	return database.Conn(&cfg.DB).Database("item_db")
}

func ItemMigrate(ctx context.Context, cfg *config.Config) {
	db := itemDBConn(cfg)
	defer db.Client().Disconnect(ctx)

	//items
	col := db.Collection("items")
	indexes, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "title", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	documents := func() []interface{} {
		items := []*item.Item{
			{
				Title:       "Diamond Sword",
				Price:       1000,
				ImageURL:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      100,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Iron Sword",
				Price:       500,
				ImageURL:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      50,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			}, {
				Title:       "Wooden Sword",
				Price:       100,
				ImageURL:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				Damage:      20,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
		}

		docs := make([]interface{}, 0)

		for _, i := range items {
			docs = append(docs, i)
		}
		return docs
	}()

	results, err := col.InsertMany(ctx, documents)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate item completed: ", results)
}
