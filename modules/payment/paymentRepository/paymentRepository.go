package paymentRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	PaymentRepository interface {
	}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) authDBConn(ctx context.Context) *mongo.Database {
	return r.db.Database("payment_db")
}
