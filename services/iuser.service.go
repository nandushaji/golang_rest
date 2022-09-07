package services

import "github.com/nandushaji/golang_rest/models"

type IUserService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(*int) (*models.User, error)
	GetAllUsers() *[]models.User
	UpdateUser(*int, *models.User) (*models.User, error)
	DeleteUser(*int) error
}
