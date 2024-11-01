package server

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
)

func (s *server) authServer() {
	repository := authRepository.NewAuthRepository(s.db)
	service := authService.NewAuthService(repository)
	httpHandler := authHandler.NewAuthHttpHandler(s.config, service)
	grpcHandler := authHandler.NewAuthGrpcHandler(service)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth/v1")

	_ = auth
}
