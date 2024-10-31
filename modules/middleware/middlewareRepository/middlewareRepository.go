package middlewareRepository

type (
	MiddlewareRepository interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepository() MiddlewareRepository {
	return &middlewareRepository{}
}
