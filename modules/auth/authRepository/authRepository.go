package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	AuthRepository interface{}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) authDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
