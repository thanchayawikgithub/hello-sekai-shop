package grpc

import (
	"log"
	"net"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	authPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authPb"
	inventoryPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryPb"
	itemPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemPb"
	playerPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.NewClient(host, opts...)
	if err != nil {
		log.Printf("failed to create grpc client: %v", err)
		return nil, err
	}

	return &grpcClientFactory{clientConn}, nil
}

func NewGrpcServer(config *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcServer, listen
}
