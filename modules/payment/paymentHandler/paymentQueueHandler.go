package paymentHandler

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/config"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentService"
)

type (
	PaymentQueueHandler interface {
	}

	paymentQueueHandler struct {
		paymentService paymentService.PaymentService
		config         *config.Config
	}
)

func NewPaymentQueueHandler(paymentService paymentService.PaymentService, config *config.Config) PaymentQueueHandler {
	return &paymentQueueHandler{paymentService, config}
}
