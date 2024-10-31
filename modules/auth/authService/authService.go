package authService

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/auth/authRepository"

type (
	AuthService interface{}

	authService struct {
		authRepo authRepository.AuthRepository
	}
)

func NewAuthService(authRepo authRepository.AuthRepository) AuthService {
	return &authService{authRepo}
}
