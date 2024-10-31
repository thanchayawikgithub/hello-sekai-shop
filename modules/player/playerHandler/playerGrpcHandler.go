package playerHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
)

type (
	playerGrpcHandler struct {
		playerService playerService.PlayerService
	}
)

func NewPlayerGrpcHandler(playerService playerService.PlayerService) *playerGrpcHandler {
	return &playerGrpcHandler{playerService}
}
