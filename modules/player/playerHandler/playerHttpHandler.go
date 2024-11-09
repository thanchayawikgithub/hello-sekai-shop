package playerHandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/custom"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/response"
)

type (
	PlayerHttpHandler interface {
		CreatePlayer(c echo.Context) error
		GetPlayerProfile(c echo.Context) error
	}

	playerHttpHandler struct {
		playerService playerService.PlayerService
		config        *config.Config
	}
)

func NewPlayerHttpHandler(playerService playerService.PlayerService, config *config.Config) PlayerHttpHandler {
	return &playerHttpHandler{playerService, config}
}

func (h *playerHttpHandler) CreatePlayer(c echo.Context) error {
	req := custom.NewCustomRequest(c)

	var playerReq player.CreatePlayerReq

	if err := req.Bind(&playerReq); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	res, err := h.playerService.CreatePlayer(c.Request().Context(), &playerReq)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusCreated, res)
}

func (h *playerHttpHandler) GetPlayerProfile(c echo.Context) error {
	playerID := strings.TrimPrefix(c.Param("player_id"), "player:")

	res, err := h.playerService.GetPlayerProfile(c.Request().Context(), playerID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return response.Success(c, http.StatusOK, res)
}
