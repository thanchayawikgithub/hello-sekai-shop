package paymentHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentService"
)

type (
	PaymentHttpHandler interface {
	}

	paymentHttpHandler struct {
		paymentService paymentService.PaymentService
		config         *config.Config
	}
)

func NewPaymentHttpHandler(paymentService paymentService.PaymentService, config *config.Config) PaymentHttpHandler {
	return &paymentHttpHandler{paymentService, config}
}
