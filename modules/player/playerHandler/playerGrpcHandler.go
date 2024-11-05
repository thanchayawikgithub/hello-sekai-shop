package playerHandler

import (
	"context"

	playerPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerService"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerService playerService.PlayerService
	}
)

func NewPlayerGrpcHandler(playerService playerService.PlayerService) *playerGrpcHandler {
	return &playerGrpcHandler{
		UnimplementedPlayerGrpcServiceServer: playerPb.UnimplementedPlayerGrpcServiceServer{},
		playerService:                        playerService,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
