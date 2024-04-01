package repository

import (
	"context"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type invoiceRepo struct {
	db *sqlx.DB
}

func (i invoiceRepo) CreateSalesInvoices(ctx context.Context, body dto.PostSalesInvoicesJSONRequestBody) (*http.Response, error) {
	// Begin a transaction
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Insert the client information
	clientQuery := "INSERT INTO clients (id, name, mobile_no, email) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, clientQuery, body.ClientId, body.Customer.Name, body.Customer.MobileNo, body.Customer.Email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if body.Items != nil {
		for _, item := range *body.Items {
			itemQuery := "INSERT INTO sales_invoice_items (product_sku, quantity, selling_price, taxed_selling_price) VALUES (?, ?, ?, ?)"
			_, err = tx.ExecContext(ctx, itemQuery, item.Product.Sku, item.Quantity, item.SellingPrice, item.TaxedSellingPrice)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	//  HTTP 201---> Created status code
	return &http.Response{
		StatusCode: http.StatusCreated,
	}, nil
}

func NewInvoiceRepo(db *sqlx.DB) adapters.InvoiceRepo {
	return &invoiceRepo{
		db: db,
	}
}
