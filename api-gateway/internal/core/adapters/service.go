package adapters

import (
	"context"
	"github.com/Nerds-Catapult/notiwa/internal/dto"
)

type AuthRepo interface {
	Login(ctx context.Context,username, password string) (*dto.Token, error)

}
