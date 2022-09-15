package services

import "github.com/golang-jwt/jwt/v4"

type IAuthenticationService interface {
	SignToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
