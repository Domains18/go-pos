package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

type productsRepo struct {
	db *sqlx.DB
}

func (p productsRepo) GetProducts(ctx context.Context, params *dto.GetProductsParams) (*http.Response, error) {
	baseQuery := "SELECT sku, name, price, tax_rate, taxed_price, image, product_category, sale_price, taxed_sale_price, quantity, available FROM products"
	var query string
	var args []interface{}

	if params.Group != nil {
		query = fmt.Sprintf("%s WHERE product_category = ?", baseQuery)
		args = append(args, *params.Group)
	} else {
		query = baseQuery
	}

	// Handle pagination
	var page int
	if params.Page != nil {
		page = *params.Page
	} else {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit
	query = fmt.Sprintf("%s LIMIT ? OFFSET ?", query)
	args = append(args, limit, offset)

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the rows into a slice of Product structs
	var products []dto.Product
	for rows.Next() {
		var product dto.Product
		if err := rows.Scan(&product.Sku, &product.Name, &product.Price, &product.TaxRate, &product.TaxedPrice, &product.Image, &product.ProductCategory, &product.SalePrice, &product.TaxedSalePrice, &product.Quantity, &product.Available); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	responseBody, err := json.Marshal(map[string]interface{}{
		"data": products,
	})
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(responseBody)),
	}, nil
}

func NewProductsRepo(db *sqlx.DB) *productsRepo {
	return &productsRepo{
		db: db,
	}
}
