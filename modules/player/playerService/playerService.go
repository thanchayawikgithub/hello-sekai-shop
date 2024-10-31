package playerService

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerRepository"
)

type (
	PlayerService interface {
	}

	playerService struct {
		playerRepo playerRepository.PlayerRepository
	}
)

func NewPlayerService(playerRepo playerRepository.PlayerRepository) PlayerService {
	return &playerService{playerRepo}
}
