package services

import (
	"context"
	"fmt"

	"github.com/nandushaji/golang_rest/models"
	"gorm.io/gorm"
)

type UserService struct {
	userDb *gorm.DB
	ctx    context.Context
}

func NewUserService(userDb *gorm.DB, ctx context.Context) IUserService {
	return &UserService{
		userDb: userDb,
		ctx:    ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	result := u.userDb.Create(&user)
	fmt.Printf(result.Name())
	return user, nil
}

func (u *UserService) GetUser(id *int) (*models.User, error) {

	return nil, nil
}

func (u *UserService) UpdateUser(id *int, user *models.User) (*models.User, error) {

	return nil, nil
}

func (u *UserService) GetAllUsers() *[]models.User {
	return nil
}
func (u *UserService) DeleteUser(id *int) error {
	return nil
}
