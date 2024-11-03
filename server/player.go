package server

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
)

func (s *server) playerServer() {
	repository := playerRepository.NewPlayerRepository(s.db)
	service := playerService.NewPlayerService(repository)
	httpHandler := playerHandler.NewPlayerHttpHandler(service, s.config)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(service)
	gueueHandler := playerHandler.NewPlayerQueueHandler(service, s.config)

	_ = httpHandler
	_ = grpcHandler
	_ = gueueHandler

	player := s.app.Group("/player/v1")
	player.GET("/health", s.healthCheckService)
}
