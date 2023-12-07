package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
)
type err *error

type ProductRepository interface {
	createProduct(ctx context.Context, product *domain.Product)(*domain.Product, error)
	GetProduct(ctx context.Context, id uint64)(*domain.Product, err)
	ListProducts(ctx context.Context, search string, categoryId, skip, limit uint64)([]domain.Product, err)
	UpdateProduct(ctx context.Context, product *domain.Product)(*domain.Product, err)
	DeleteProduct(ctx context.Context, id uint64) err
}