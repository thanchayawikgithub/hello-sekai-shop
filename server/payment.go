package server

import (
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentHandler"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentRepository"
	"github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentService"
)

func (s *server) paymentServer() {
	repository := paymentRepository.NewPaymentRepository(s.db)
	service := paymentService.NewPaymentService(repository)
	httpHandler := paymentHandler.NewPaymentHttpHandler(service, s.config)
	queueHandler := paymentHandler.NewPaymentQueueHandler(service, s.config)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")
	payment.GET("", s.healthCheckService)
}
