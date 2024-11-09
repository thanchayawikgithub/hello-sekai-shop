package item

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	Item struct {
		ID          bson.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string        `json:"title" bson:"title"`
		ImageURL    string        `json:"image_url" bson:"image_url"`
		Price       float64       `json:"price" bson:"price"`
		Damage      int           `json:"damage" bson:"damage"`
		UsageStatus bool          `json:"usage_status" bson:"usage_status"`
		CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	}
)
