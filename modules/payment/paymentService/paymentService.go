package paymentService

import "github.com/thanchayawikgithub/hello-sekai-shop/modules/payment/paymentRepository"

type (
	PaymentService interface {
	}

	paymentService struct {
		paymentRepo paymentRepository.PaymentRepository
	}
)

func NewPaymentService(paymentRepo paymentRepository.PaymentRepository) PaymentService {
	return &paymentService{paymentRepo}
}
