package repository

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type loginRepo struct {
	db *sqlx.DB

}

func (l loginRepo) LoginCustomer(ctx context.Context, body dto.PostLoginJSONRequestBody) (*http.Response, error) {
	var userID int
	var displayName string
	query := "SELECT id, display_name FROM users WHERE username = ? AND password = ?"
	err := l.db.QueryRowContext(ctx, query, body.Username, body.Password).Scan(&userID, &displayName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Credentials are incorrect
			return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       ioutil.NopCloser(strings.NewReader(`{"code":"401", "message":"Invalid credentials"}`)),
			}, nil
		}
		// Other error
		return nil, err
	}

	// Generate a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	})
	tokenString, err := token.SignedString([]byte("your_secret_key")) // Replace with your secret key
	if err != nil {
		return nil, err
	}

	// Construct the response body
	responseBody, err := json.Marshal(map[string]interface{}{
		"token":        tokenString,
		"user":         userID,
		"expire":       time.Now().Add(time.Hour * 24).Unix(),
		"display_name": displayName,
	})
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(responseBody)),
	}, nil
}

func NewLoginRepo(db *sqlx.DB) adapters.LoginRepo{
	return &loginRepo{
		db: db,
	}
}