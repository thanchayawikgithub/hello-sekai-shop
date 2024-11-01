package server

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemService"
)

func (s *server) itemServer() {
	repository := itemRepository.NewItemRepository(s.db)
	service := itemService.NewItemService(repository)
	httpHandler := itemHandler.NewItemHttpHandler(service, s.config)
	grpcHandler := itemHandler.NewItemGrpcHandler(service)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item/v1")

	_ = item
}
