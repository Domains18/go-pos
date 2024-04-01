package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/core/adapters"
	"github.com/Nerds-Catapult/notiwa/api-gateway/internal/dto"
	"github.com/Nerds-Catapult/notiwa/api-gateway/pkg/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
)

type authRepo struct {
	db *sqlx.DB
}

func (a *authRepo) Login(ctx context.Context,username, password string) (*dto.Token, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	var user struct {
		ID           string `db:"user_id"`
		PasswordHash string `db:"password_hash"`
		DisplayName  string `db:"display_name"`
	}

	err := a.db.Get(&user, "SELECT user_id, password_hash, display_name FROM users WHERE username = $1", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	generatedJwtToken, err := utils.GenerateJWT(username)
	if err != nil {
		return nil, err
	}

	token := &dto.Token{
		Token:       &generatedJwtToken,
		User:        &username,
		//Expire:      ,
		DisplayName: &user.DisplayName,
	}

	return token, nil
}


func NewAuthRepo(db *sqlx.DB) adapters.AuthRepo{
	return &authRepo{db: db}
}