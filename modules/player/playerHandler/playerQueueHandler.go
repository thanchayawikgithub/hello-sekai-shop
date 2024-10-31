package playerHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
)

type (
	PlayerQueueHandler interface {
	}

	playerQueueHandler struct {
		playerService playerService.PlayerService
		config        *config.Config
	}
)

func NewPlayerQueueHandler(playerService playerService.PlayerService, config *config.Config) PlayerQueueHandler {
	return &playerQueueHandler{playerService, config}
}
