package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	GetPaymentById(ctx context.Context, id uint64) (*domain.Payment, error)
}