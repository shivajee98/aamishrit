package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) GetPaymentByTransactionID(transactionID string) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Where("transaction_id = ?", transactionID).First(&payment).Error
	return &payment, err
}

func (r *PaymentRepository) UpdatePaymentStatus(transactionID, status string) error {
	return r.db.Model(&model.Payment{}).Where("transaction_id = ?", transactionID).Update("status", status).Error
}
