package port

import (
	"errors"
	"strings"
)

var (
	ErrDataNotFound               = errors.New("data not found")
	ErrNoUpdateData               = errors.New("no data to update")
	ErrConflictingData            = errors.New("data conflicts with existing data in unique column")
	ErrInsufficientStock          = errors.New("product stock is not enough")
	ErrInsufficientPayment        = errors.New("insufficient payment")
	ErrExpiredToken               = errors.New("access token has expired")
	ErrInvalidCredentials         = errors.New("invalid email or password")
	ErrEmptyAuthorizationHeader   = errors.New("no authorization headers were provided")
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	ErrInvalidAuthorizationType   = errors.New("authorization type is not supported")
	ErrUnauthorized               = errors.New("user is unauthorized to access the resource")
	ErrForbidden                  = errors.New("user is forbidden to access the resource")
)

func isUniqueConstraintViolation(err error) bool {
	return strings.Contains(err.Error(), "23505")
}
