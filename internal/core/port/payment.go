package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	GetPaymentById(ctx context.Context, id uint64) (*domain.Payment, error)
	ListPayment(ctx context.Context, skip, limit uint64) (*domain.Payment, error)
	UpdatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	DeletePayment(ctx context.Context, id uint64) error
}

type PaymentService interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	GetPayment(ctx context.Context, id uint64) (*domain.Payment, error)
	ListPayments(ctx context.Context, skip, limit uint64)
	UpdatePayment(ctx context.Context, id uint64) error
}
