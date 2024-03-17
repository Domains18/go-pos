package adapters

import (
	"context"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"net/http"
)
type AuthRepo interface {
	Login(ctx context.Context, username, password string) (*dto.Token, error)
}

type Server interface {
	CreateCustomers(ctx context.Context, body dto.PostCustomersJSONRequestBody) (*http.Response, error)
	GetCustomers(ctx context.Context, params *dto.GetCustomersParams) (*http.Response, error)
	LoginCus(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error)
	GetProducts(ctx context.Context, params *dto.GetProductsParams) (*http.Response, error)
	CreateSalesInvoices(ctx context.Context, body dto.PostSalesInvoicesJSONRequestBody) (*http.Response, error)
}
