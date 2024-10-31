package middlewareService

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/middleware/middlewareRepository"

type (
	MiddlewareService interface{}

	middlewareService struct {
		middlewareRepo middlewareRepository.MiddlewareRepository
	}
)

func NewMiddlewareService(middlewareRepo middlewareRepository.MiddlewareRepository) MiddlewareService {
	return &middlewareService{middlewareRepo}
}
