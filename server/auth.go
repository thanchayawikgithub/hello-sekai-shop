package server

import (
	"log"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authHandler"
	authPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/grpc"
)

func (s *server) authServer() {
	repository := authRepository.NewAuthRepository(s.db)
	service := authService.NewAuthService(repository)
	httpHandler := authHandler.NewAuthHttpHandler(s.config, service)
	grpcHandler := authHandler.NewAuthGrpcHandler(service)

	//Grpc
	go func() {
		gprcServer, listen := grpc.NewGrpcServer(&s.config.Jwt, s.config.Grpc.AuthURL)
		authPb.RegisterAuthGrpcServiceServer(gprcServer, grpcHandler)

		gprcServer.Serve(listen)
		log.Println("Grpc server is running on", s.config.Grpc.AuthURL)
	}()

	_ = grpcHandler

	auth := s.app.Group("/auth_v1")
	auth.GET("", s.healthCheckService)

	auth.POST("/login", httpHandler.Login)
}
