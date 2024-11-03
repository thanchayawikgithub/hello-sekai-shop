package response

import "github.com/labstack/echo/v4"

type (
	ErrorMessage struct {
		Message string `json:"message"`
	}

	CustomResponse interface {
		Error(code int, err error) error
		Success(code int, data interface{}) error
	}
)

func Error(ctx echo.Context, code int, err error) error {
	return ctx.JSON(code, &ErrorMessage{Message: err.Error()})
}

func Success(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(code, data)
}
