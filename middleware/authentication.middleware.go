package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nandushaji/golang_rest/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("uuid", uuid)
		fmt.Printf("The request with uuid %s is started \n", uuid)
		token := c.GetHeader("Authorization")
		if t, err := services.NewAuthenticationService().ValidateToken(token); err == nil {
			fmt.Println(t.Claims.(*jwt.MapClaims))
			c.Next()
		}
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}
