package server

import (
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemHandler"
	itemPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/item/itemService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/grpc"
)

func (s *server) itemServer() {
	repository := itemRepository.NewItemRepository(s.db)
	service := itemService.NewItemService(repository)
	httpHandler := itemHandler.NewItemHttpHandler(service, s.config)
	grpcHandler := itemHandler.NewItemGrpcHandler(service)

	//Grpc
	go func() {
		gprcServer, listen := grpc.NewGrpcServer(&s.config.Jwt, s.config.Grpc.ItemURL)
		itemPb.RegisterItemGrpcServiceServer(gprcServer, grpcHandler)

		gprcServer.Serve(listen)
		log.Println("Grpc server is running on", s.config.Grpc.ItemURL)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")
	item.GET("", s.healthCheckService)
}
