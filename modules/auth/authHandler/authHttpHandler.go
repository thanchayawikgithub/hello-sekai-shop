package authHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
)

type (
	AuthHttpHandler interface{}

	authHttpHandler struct {
		authService authService.AuthService
		config      *config.Config
	}
)

func NewAuthHttpHandler(config *config.Config, authService authService.AuthService) AuthHttpHandler {
	return &authHttpHandler{authService, config}
}
