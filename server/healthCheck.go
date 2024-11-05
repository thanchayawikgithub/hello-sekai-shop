package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/response"
)

type healthCheck struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func (s *server) healthCheckService(ctx echo.Context) error {
	return response.Success(ctx, http.StatusOK, &healthCheck{App: s.config.App.Name, Status: "OK"})
}
