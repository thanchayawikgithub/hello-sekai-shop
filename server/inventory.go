package server

import (
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryHandler"
	inventoryPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/grpc"
)

func (s *server) inventoryServer() {
	repository := inventoryRepository.NewInventoryRepository(s.db)
	service := inventoryService.NewInventoryService(repository)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(service, s.config)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(service)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(service, s.config)

	//Grpc
	go func() {
		gprcServer, listen := grpc.NewGrpcServer(&s.config.Jwt, s.config.Grpc.InventoryURL)
		inventoryPb.RegisterInventoryGrpcServiceServer(gprcServer, grpcHandler)

		gprcServer.Serve(listen)
		log.Println("Grpc server is running on", s.config.Grpc.InventoryURL)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")
	inventory.GET("", s.healthCheckService)
}
