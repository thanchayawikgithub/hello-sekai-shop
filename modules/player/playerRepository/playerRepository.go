package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	PlayerRepository interface {
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	return &playerRepository{db}
}

func (r *playerRepository) authDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
