package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategoryByID(ctx context.Context, id uint64) (*domain.Category, error)
	ListCategories(ctx context.Context, skip, limit uint64) ([]domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id uint64) error
}

type categoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategory(ctx context.Context, id uint64) (*domain.Category, error)
	ListCategories(ctx context.Context, skip, limit uint64) ([]domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id uint64) error
}
