package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nandushaji/golang_rest/models"
	"github.com/nandushaji/golang_rest/services"
)

type UserController struct {
	UserService services.IUserService
}

func NewUserController(userService services.IUserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	newUser, err := u.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (u *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Print(id)
	user, err := u.UserService.GetUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	updatedUser, err := u.UserService.UpdateUser(&id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	users := u.UserService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := u.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success!"})
}

func (u *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", u.CreateUser)
	userroute.GET("/", u.GetAllUsers)
	userroute.GET("/:id", u.GetUser)
	userroute.PUT("/:id", u.UpdateUser)
	userroute.DELETE("/:id", u.DeleteUser)
}
