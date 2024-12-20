package authHandler

import (
	"context"

	authPb "github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authPb"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authService"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authService authService.AuthService
	}
)

func NewAuthGrpcHandler(authService authService.AuthService) *authGrpcHandler {
	return &authGrpcHandler{
		UnimplementedAuthGrpcServiceServer: authPb.UnimplementedAuthGrpcServiceServer{},
		authService:                        authService,
	}
}

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
