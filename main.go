package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nandushaji/golang_rest/config"
	"github.com/nandushaji/golang_rest/controllers"
	"github.com/nandushaji/golang_rest/services"
	"gorm.io/gorm"
)

var (
	server                *gin.Engine
	userService           services.IUserService
	userController        controllers.UserController
	ctx                   context.Context
	db                    *gorm.DB
	loginService          services.ILoginService
	authenticationService services.IAuthenticationService
	loginController       controllers.LoginController
)

func init() {
	ctx = context.TODO()
	db, _ = config.Connect()
	userService = services.NewUserService(db, ctx)
	userController = controllers.NewUserController(userService)
	authenticationService = services.NewAuthenticationService()
	loginService = services.NewLoginService(ctx, authenticationService)
	loginController = controllers.NewLoginController(loginService)
	server = gin.Default()
}

func main() {
	basepath := server.Group("api/v1")
	userController.RegisterUserRoutes(basepath)
	loginController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":3000"))
}
