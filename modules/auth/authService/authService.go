package authService

import (
	"context"
	"time"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/player"
	playerPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/player/playerPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/jwtauth"
	"github.com/thanchayawikgithub/hello-sekai-shop/pkg/utils"
)

type (
	AuthService interface {
		Login(ctx context.Context, cfg *config.Config, req *auth.PlayerLoginReq) (*auth.ProfileInterceptor, error)
	}

	authService struct {
		authRepo authRepository.AuthRepository
	}
)

func NewAuthService(authRepo authRepository.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Login(ctx context.Context, cfg *config.Config, req *auth.PlayerLoginReq) (*auth.ProfileInterceptor, error) {
	profile, err := s.authRepo.CredentialSearch(ctx, cfg.Grpc.PlayerURL, &playerPb.CredentialSearchReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	profile.Id = "player:" + profile.Id
	accessToken := jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		PlayerID: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtauth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtauth.Claims{
		PlayerID: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	credentialID, err := s.authRepo.InsertCredential(ctx, &auth.Credential{
		PlayerID:     profile.Id,
		RoleCode:     int(profile.RoleCode),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    utils.LocalTime(),
	})

	loc, _ := time.LoadLocation("Asia/Bangkok")
	return &auth.ProfileInterceptor{
		PlayerProfile: &player.PlayerProfile{
			ID:        profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringToTime(profile.CreatedAt).In(loc),
			UpdatedAt: utils.ConvertStringToTime(profile.UpdatedAt).In(loc),
		},
		Credential: &auth.CredentialRes{
			ID:           credentialID.Hex(),
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
