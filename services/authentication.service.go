package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticationClaims struct {
	email  string
	isUser bool
	jwt.StandardClaims
}

type AuthenticationService struct {
	secret string
}

func NewAuthenticationService() IAuthenticationService {
	return &AuthenticationService{secret: "hgdhgdgfwdegwyhdfew"}
}

func (a *AuthenticationService) SignToken(email string, isUser bool) string {
	claims := &AuthenticationClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "X_X",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(a.secret))
	fmt.Println(t, err)
	if err == nil {
		return t
	}
	return ""
}

func (a *AuthenticationService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			fmt.Println(claims)
		}
		return []byte(a.secret), nil
	})

}
