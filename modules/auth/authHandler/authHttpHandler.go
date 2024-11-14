package authHandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/custom"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/response"
)

type (
	AuthHttpHandler interface {
		Login(c echo.Context) error
	}

	authHttpHandler struct {
		authService authService.AuthService
		config      *config.Config
	}
)

func NewAuthHttpHandler(config *config.Config, authService authService.AuthService) AuthHttpHandler {
	return &authHttpHandler{authService, config}
}

func (h *authHttpHandler) Login(c echo.Context) error {
	ctx := context.Background()

	customReq := custom.NewCustomRequest(c)

	req := new(auth.PlayerLoginReq)
	if err := customReq.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	res, err := h.authService.Login(ctx, h.config, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, res)
}
