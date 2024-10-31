package inventoryRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	InventoryRepository interface{}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) inventoryDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("inventory_db")
}
