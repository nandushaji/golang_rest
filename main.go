package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nandushaji/golang_rest/controllers"
	"github.com/nandushaji/golang_rest/models"
	"github.com/nandushaji/golang_rest/services"
)

var (
	server         *gin.Engine
	userService    services.IUserService
	userController controllers.UserController
	ctx            context.Context
)

func init() {
	ctx = context.TODO()
	userCollection := []models.User{}
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	basepath := server.Group("api/v1")
	userController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":3000"))
}
