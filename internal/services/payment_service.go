package services

import (
	"errors"
	"os"

	"github.com/razorpay/razorpay-go"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type PaymentService struct {
	paymentRepo *repository.PaymentRepository
	razorpay    *razorpay.Client
}

func NewPaymentService(paymentRepo *repository.PaymentRepository) *PaymentService {
	client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY"), os.Getenv("RAZORPAY_SECRET"))
	return &PaymentService{
		paymentRepo: paymentRepo,
		razorpay:    client,
	}
}

func (s *PaymentService) CreateOrder(amount float64, userID, orderID uint) (*model.Payment, error) {
	data := map[string]interface{}{
		"amount":   int(amount * 100), // Razorpay expects amount in paise
		"currency": "INR",
	}

	order, err := s.razorpay.Order.Create(data, nil)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{
		UserID:          userID,
		OrderID:         orderID,
		Amount:          amount,
		Currency:        "INR",
		Status:          "pending",
		RazorpayOrderID: order["id"].(string),
	}

	err = s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) VerifyPayment(transactionID, razorpayOrderID string) error {
	payment, err := s.paymentRepo.GetPaymentByTransactionID(transactionID)
	if err != nil {
		return errors.New("payment not found")
	}

	if payment.RazorpayOrderID != razorpayOrderID {
		return errors.New("invalid order ID for this transaction")
	}

	err = s.paymentRepo.UpdatePaymentStatus(transactionID, "success")
	return err
}
