package middlewareHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/middleware/middlewareService"
)

type (
	MiddlewareHandler interface{}

	middlewareHandler struct {
		middlewareService middlewareService.MiddlewareService
		config            *config.Config
	}
)

func NewMiddlewareHandler(middlewareService middlewareService.MiddlewareService, config *config.Config) MiddlewareHandler {
	return &middlewareHandler{middlewareService, config}
}
