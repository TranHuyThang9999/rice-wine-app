package apperrors

import "errors"

var (
	ErrUserNotFound            = errors.New("user not found")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenExpired            = errors.New("token has expired")
	ErrInvalidExpiration       = errors.New("invalid expiration in token")
)
