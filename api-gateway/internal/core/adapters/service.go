package adapters

import (
	"context"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"net/http"
)
type AuthRepo interface {
	Login(ctx context.Context, username, password string) (*dto.Token, error)
}

type CustomerRepo interface {
	CreateCustomers(ctx context.Context, body dto.PostCustomersJSONRequestBody) (*http.Response, error)
	GetCustomers(ctx context.Context, params *dto.GetCustomersParams) (*http.Response, error)
}
type ProductRepo interface {
	GetProducts(ctx context.Context, params *dto.GetProductsParams) (*http.Response, error)

}

type InvoiceRepo interface {
	CreateSalesInvoices(ctx context.Context, body dto.PostSalesInvoicesJSONRequestBody) (*http.Response, error)
}

type LoginRepo interface {
	LoginCustomer(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error)
}


type Server interface {
	CreateCustomers(ctx context.Context, body dto.PostCustomersJSONRequestBody) (*http.Response, error)
	GetCustomers(ctx context.Context, params *dto.GetCustomersParams) (*http.Response, error)
	LoginCustomer(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error)
	GetProducts(ctx context.Context, params *dto.GetProductsParams) (*http.Response, error)
	CreateSalesInvoices(ctx context.Context, body dto.PostSalesInvoicesJSONRequestBody) (*http.Response, error)
}

