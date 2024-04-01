package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"net/http"
)

type AuthHandler struct {
	AuthRepo adapters.AuthRepo
}

func (h *AuthHandler) CreateCustomers(ctx context.Context, body dto.PostCustomersJSONRequestBody) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) GetCustomers(ctx context.Context, params *dto.GetCustomersParams) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) LoginCustomer(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) GetProducts(ctx context.Context, params *dto.GetProductsParams) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) CreateSalesInvoices(ctx context.Context, body dto.PostSalesInvoicesJSONRequestBody) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}



func NewAuthHandler(authRepo adapters.AuthRepo) *AuthHandler {
	return &AuthHandler{
		AuthRepo: authRepo,
	}
}