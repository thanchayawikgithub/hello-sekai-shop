package playerRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	PlayerRepository interface {
		IsUniquePlayer(ctx context.Context, email, username string) bool
		InsertOnePlayer(ctx context.Context, req *player.Player) (bson.ObjectID, error)
		FindOnePlayerProfile(ctx context.Context, playerID string) (*player.PlayerProfileBson, error)
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	return &playerRepository{db}
}

func (r *playerRepository) playerDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}

func (r *playerRepository) IsUniquePlayer(ctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db := r.playerDBConn(ctx)
	col := db.Collection("players")

	player := new(player.Player)

	if err := col.FindOne(
		ctx,
		bson.M{
			"$or": []bson.M{
				{"username": username},
				{"email": email},
			},
		},
	).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %v", err)
		return true
	}

	return false
}

func (r *playerRepository) InsertOnePlayer(ctx context.Context, req *player.Player) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db := r.playerDBConn(ctx)
	col := db.Collection("players")

	playerID, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %v", err)
		return bson.NilObjectID, errors.New("error: insert one player failed")
	}

	return playerID.InsertedID.(bson.ObjectID), nil
}

func (r *playerRepository) FindOnePlayerProfile(ctx context.Context, playerID string) (*player.PlayerProfileBson, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db := r.playerDBConn(ctx)
	col := db.Collection("players")

	result := new(player.PlayerProfileBson)

	playerObjID, err := utils.ConvertToObject(playerID)
	if err != nil {
		log.Printf("Error: FindOnePlayerProfile: %v", err)
		return nil, errors.New("error: invalid player id")
	}

	if err := col.FindOne(ctx, bson.M{"_id": playerObjID}, options.FindOne().SetProjection((bson.M{"_id": 1, "email": 1, "username": 1, "created_at": 1, "updated_at": 1}))).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerProfile: %v", err)
		return nil, errors.New("error: find one player profile failed")
	}

	return result, nil
}
