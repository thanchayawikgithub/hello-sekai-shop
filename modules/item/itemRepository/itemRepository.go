package itemRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	ItemRepository interface{}

	itemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) itemDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
