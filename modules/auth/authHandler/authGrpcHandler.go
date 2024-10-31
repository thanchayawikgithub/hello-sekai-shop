package authHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
)

type (
	authGrpcHandler struct {
		authService authService.AuthService
	}
)

func NewAuthGrpcHandler(authService authService.AuthService) *authGrpcHandler {
	return &authGrpcHandler{authService}
}
