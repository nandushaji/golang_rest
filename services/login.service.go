package services

import (
	"context"
)

type LoginService struct {
	ctx                   context.Context
	AuthenticationService IAuthenticationService
}

func NewLoginService(ctx context.Context, authenticationService IAuthenticationService) ILoginService {
	return &LoginService{
		ctx:                   ctx,
		AuthenticationService: authenticationService,
	}
}

func (l *LoginService) AdminLogin(email, password string) string {
	adminEmail := "admin@example.com"
	adminPassword := "1234"
	if adminEmail == email && adminPassword == password {
		token := l.AuthenticationService.SignToken(adminEmail, true)
		return token
	}
	return ""
}
