package handlers

import (
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
)

type Handler struct {
	AuthRepo adapters.AuthRepo
	ServRepo adapters.Server

	CustomerRepo adapters.CustomerRepo
	ProductRepo  adapters.ProductRepo
	InvoiceRepo  adapters.InvoiceRepo
	LoginRepo    adapters.LoginRepo
}

func NewHandler(authRepo adapters.AuthRepo, servRepo adapters.Server, CustomerRepo adapters.CustomerRepo, ProductRepo adapters.ProductRepo) *Handler {
	return &Handler{
		AuthRepo: authRepo,
		ServRepo: servRepo,
	}
}

//
//func (h Handler) LoginCustomer(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error) {
//	return h.LoginRepo.LoginCustomer(ctx, body)
//}
