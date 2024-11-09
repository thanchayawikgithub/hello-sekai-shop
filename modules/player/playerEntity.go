package player

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	Player struct {
		ID          bson.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email       string        `json:"email" bson:"email"`
		Password    string        `json:"password" bson:"password"`
		Username    string        `json:"username" bson:"username"`
		CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
		PlayerRoles []PlayerRole  `json:"player_roles"`
	}

	PlayerRole struct {
		RoleTitle string `json:"role_title" bson:"role_title"`
		RoleCode  int    `json:"role_code" bson:"role_code"`
	}

	PlayerProfileBson struct {
		ID        bson.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email     string        `json:"email" bson:"email"`
		Username  string        `json:"username" bson:"username"`
		CreatedAt time.Time     `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	}

	PlayerSavingAccount struct {
		PlayerID string  `json:"player_id" bson:"player_id"`
		Balance  float64 `json:"balance" bson:"balance"`
	}

	PlayerTransaction struct {
		PlayerID  string    `bson:"player_id"`
		Amount    float64   `bson:"amount"`
		CreatedAt time.Time `bson:"created_at"`
	}
)
