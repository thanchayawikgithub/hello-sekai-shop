package migration

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func authDBConn(cfg *config.Config) *mongo.Database {
	return database.Conn(&cfg.DB).Database("auth_db")
}

func AuthMigrate(ctx context.Context, cfg *config.Config) {
	db := authDBConn(cfg)
	defer db.Client().Disconnect(ctx)

	//auth
	col := db.Collection("auth")
	indexes, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
		{Keys: bson.D{{Key: "refresh_token", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	//roles
	col = db.Collection("roles")
	indexes, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "code", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	documents := func() []interface{} {
		roles := []*auth.Role{
			{
				Title: "player",
				Code:  0,
			},
			{
				Title: "admin",
				Code:  1,
			},
		}

		docs := make([]interface{}, 0)

		for _, role := range roles {
			docs = append(docs, role)
		}
		return docs
	}()

	results, err := col.InsertMany(ctx, documents)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate auth completed: ", results)
}
