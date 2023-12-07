package port

import (
	"context"
	"github.com/Domains18/go-pos/internal/core/domain"
	"time"
)

type TokenService interface {
	CreateToken(user *domain.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
