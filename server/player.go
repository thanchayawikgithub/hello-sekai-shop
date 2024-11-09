package server

import (
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerHandler"
	playerPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/grpc"
)

func (s *server) playerServer() {
	repository := playerRepository.NewPlayerRepository(s.db)
	service := playerService.NewPlayerService(repository)
	httpHandler := playerHandler.NewPlayerHttpHandler(service, s.config)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(service)
	gueueHandler := playerHandler.NewPlayerQueueHandler(service, s.config)

	//Grpc
	go func() {
		gprcServer, listen := grpc.NewGrpcServer(&s.config.Jwt, s.config.Grpc.PlayerURL)
		playerPb.RegisterPlayerGrpcServiceServer(gprcServer, grpcHandler)

		gprcServer.Serve(listen)
		log.Println("Grpc server is running on", s.config.Grpc.PlayerURL)
	}()

	_ = grpcHandler
	_ = gueueHandler

	player := s.app.Group("/player_v1")
	player.GET("", s.healthCheckService)
	player.POST("/register", httpHandler.CreatePlayer)
	player.GET("/:player_id", httpHandler.GetPlayerProfile)
	player.POST("/transaction", httpHandler.CreatePlayerTransaction)
	player.GET("/:player_id/account", httpHandler.GetPlayerSavingAccount)
}
