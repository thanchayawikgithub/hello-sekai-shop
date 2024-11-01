package server

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/inventory/inventoryService"
)

func (s *server) inventoryServer() {
	repository := inventoryRepository.NewInventoryRepository(s.db)
	service := inventoryService.NewInventoryService(repository)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(service, s.config)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(service)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(service, s.config)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory/v1")

	_ = inventory
}
