package migration

import (
	"context"
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/database"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func playerDBConn(cfg *config.Config) *mongo.Database {
	return database.Conn(&cfg.DB).Database("player_db")
}

func PlayerMigrate(ctx context.Context, cfg *config.Config) {
	db := playerDBConn(cfg)
	defer db.Client().Disconnect(ctx)

	//player_transactions
	col := db.Collection(database.PlayerTransactionCollection)
	indexes, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	col = db.Collection(database.PlayerCollection)
	indexes, err = col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "email", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}

	for _, index := range indexes {
		log.Printf("Index: %s", index)
	}

	documents := func() []interface{} {
		players := []*player.Player{
			{
				Email:    "test@test.com",
				Password: "12345678",
				Username: "test",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "than@mail.com",
				Password: "than11014",
				Username: "than",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		docs := make([]interface{}, 0)

		for _, p := range players {
			docs = append(docs, p)
		}
		return docs
	}()

	results, err := col.InsertMany(ctx, documents)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate player completed: ", results)

	playerTransactions := make([]interface{}, 0)
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, &player.PlayerTransaction{
			PlayerID:  "player:" + p.(bson.ObjectID).Hex(),
			Amount:    10000,
			CreatedAt: utils.LocalTime(),
		})
	}

	col = db.Collection(database.PlayerTransactionCollection)
	results, err = col.InsertMany(ctx, playerTransactions)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player transactions completed: ", results)

	col = db.Collection(database.PlayerTransactionQueueCollection)
	result, err := col.InsertOne(ctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player transactions queue completed: ", result)
}
