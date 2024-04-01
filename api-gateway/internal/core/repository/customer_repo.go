package repository

import (
	"context"
	"errors"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type customerRepo struct {
	db *sqlx.DB
}

func (c customerRepo) CreateCustomers(ctx context.Context, body dto.PostCustomersJSONRequestBody) (*http.Response, error) {
	query := "INSERT INTO customers (name, mobile_no, email, id) VALUES (?, ?, ?, ?)"
	_, err := c.db.ExecContext(ctx, query, body.Name, body.MobileNo, body.Email, body.ID)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusCreated,
	}, nil
}

func (c customerRepo) GetCustomers(ctx context.Context, params *dto.GetCustomersParams) (*http.Response, error) {
	if params.Search == nil {
		return nil, errors.New("search parameter is missing")
	}
	searchStr := *params.Search // Dereference the pointer to get the string value

	query := "SELECT name, mobile_no, email, id FROM customers WHERE mobile_no LIKE ? OR id LIKE ?"
	rows, err := c.db.QueryContext(ctx, query, "%"+searchStr+"%", "%"+searchStr+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//response
	var customers []dto.Customer
	for rows.Next() {
		var customer dto.Customer
		if err := rows.Scan(&customer.Name, &customer.MobileNo, &customer.Email, &customer.ID); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		// Body should contain the JSON marshaled `customers` data
	}, nil
}

func NewCustomerRepo(db *sqlx.DB) *customerRepo {
	return &customerRepo{
		db: db,
	}
}
