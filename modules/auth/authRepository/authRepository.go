package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth"
	playerPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/grpc"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	AuthRepository interface {
		CredentialSearch(ctx context.Context, grpcURL string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error)
		InsertCredential(ctx context.Context, req *auth.Credential) (bson.ObjectID, error)
	}

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

func (r *authRepository) CredentialSearch(ctx context.Context, grpcURL string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	conn, err := grpc.NewGrpcClient(grpcURL)
	if err != nil {
		log.Println("error: create grpc client", err)
		return nil, errors.New("error: create grpc client")
	}

	result, err := conn.Player().CredentialSearch(ctx, req)
	if err != nil {
		log.Println("error: credential search failed", err.Error())
		return nil, errors.New(err.Error())
	}

	return result, nil
}

func (r *authRepository) InsertCredential(ctx context.Context, req *auth.Credential) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	db := r.authDBConn(ctx)
	col := db.Collection("auth")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("error: insert credential failed: %v", err.Error())
		return bson.NilObjectID, errors.New("error: insert credential failed")
	}

	return result.InsertedID.(bson.ObjectID), nil
}
