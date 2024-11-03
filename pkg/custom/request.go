package custom

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	CustomRequest interface {
		Bind(i interface{}) error
	}

	customRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

func NewCustomRequest(ctx echo.Context) CustomRequest {
	return &customRequest{ctx: ctx, validator: validator.New()}
}

func (r *customRequest) Bind(i interface{}) error {
	if err := r.ctx.Bind(i); err != nil {
		log.Printf("Error: Binding data failed: %s", err.Error())
		return r.ctx.JSON(http.StatusBadRequest, "Invalid request data")
	}

	if err := r.validator.Struct(i); err != nil {
		log.Printf("Error: Validating data failed:  %s", err.Error())
		return formatValidationError(r.ctx, err)
	}

	return nil
}

func formatValidationError(ctx echo.Context, err error) error {
	errs := err.(validator.ValidationErrors)
	validationErrors := make(map[string]string)
	for _, e := range errs {
		validationErrors[e.Field()] = "Validation failed on the '" + e.Tag() + "' rule"
	}
	return ctx.JSON(http.StatusBadRequest, validationErrors)
}
