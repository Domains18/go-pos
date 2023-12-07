package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uint64) (*domain.Order, error)
	ListOrders(ctx context.Context, skip, limit uint64) ([]domain.Order, error)
}

type OrderService interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrder(ctx context.Context, id uint64) (*domain.Order, error)
	ListOrder(ctx context.Context, skip, limit uint64) ([]domain.Order, error)
}
