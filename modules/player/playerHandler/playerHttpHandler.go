package playerHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
)

type (
	PlayerHttpHandler interface {
	}

	playerHttpHandler struct {
		playerService playerService.PlayerService
		config        *config.Config
	}
)

func NewPlayerHttpHandler(playerService playerService.PlayerService, config *config.Config) PlayerHttpHandler {
	return &playerHttpHandler{playerService, config}
}
