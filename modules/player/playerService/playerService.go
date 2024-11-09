package playerService

import (
	"context"
	"errors"
	"time"

	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	PlayerService interface {
		CreatePlayer(ctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error)
		GetPlayerProfile(ctx context.Context, playerID string) (*player.PlayerProfile, error)
		CreatePlayerTransaction(ctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error)
		GetPlayerSavingAccount(ctx context.Context, playerID string) (*player.PlayerSavingAccount, error)
	}

	playerService struct {
		playerRepo playerRepository.PlayerRepository
	}
)

func NewPlayerService(playerRepo playerRepository.PlayerRepository) PlayerService {
	return &playerService{playerRepo}
}

func (s *playerService) CreatePlayer(ctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error) {
	if !s.playerRepo.IsUniquePlayer(ctx, req.Email, req.Username) {
		return nil, errors.New("error: email or username already exists")
	}

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: failed to hash password")
	}

	//Insert player
	playerID, err := s.playerRepo.InsertOnePlayer(ctx, &player.Player{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
		PlayerRoles: []player.PlayerRole{
			{
				RoleTitle: "player",
				RoleCode:  0,
			},
		},
	})
	if err != nil {
		return nil, errors.New("error: failed to insert player")
	}

	return s.GetPlayerProfile(ctx, playerID.Hex())
}

func (s *playerService) GetPlayerProfile(ctx context.Context, playerID string) (*player.PlayerProfile, error) {
	result, err := s.playerRepo.FindOnePlayerProfile(ctx, playerID)
	if err != nil {
		return nil, errors.New("error: failed to get player profile")
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, errors.New("error: failed to load location")
	}

	return &player.PlayerProfile{
		ID:        result.ID.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		CreatedAt: result.CreatedAt.In(loc),
		UpdatedAt: result.UpdatedAt.In(loc),
	}, nil
}

func (s *playerService) CreatePlayerTransaction(ctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error) {
	if err := s.playerRepo.InsertOnePlayerTransaction(ctx, &player.PlayerTransaction{
		PlayerID:  req.PlayerID,
		Amount:    req.Amount,
		CreatedAt: utils.LocalTime(),
	}); err != nil {
		return nil, errors.New("error: failed to create player transaction")
	}

	return s.GetPlayerSavingAccount(ctx, req.PlayerID)
}

func (s *playerService) GetPlayerSavingAccount(ctx context.Context, playerID string) (*player.PlayerSavingAccount, error) {
	return s.playerRepo.GetPlayerSavingAccount(ctx, playerID)
}
