package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nandushaji/golang_rest/services"
)

type LoginController struct {
	LoginService services.ILoginService
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginController(loginService services.ILoginService) LoginController {
	return LoginController{
		LoginService: loginService,
	}
}

func (l *LoginController) Login(ctx *gin.Context) {
	var loginCreds LoginModel
	if err := ctx.ShouldBindJSON(&loginCreds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	token := l.LoginService.AdminLogin(loginCreds.Email, loginCreds.Password)
	if len(token) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to Authenticate"})
	}
	ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
}

func (l *LoginController) RegisterUserRoutes(rg *gin.RouterGroup) {
	loginRoute := rg.Group("/login")
	loginRoute.POST("/", l.Login)
}
