package services

import (
	"context"

	"github.com/nandushaji/golang_rest/models"
)

type UserService struct {
	userCollection []models.User
	ctx            context.Context
}

func NewUserService(userCollection []models.User, ctx context.Context) IUserService {
	return &UserService{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	u.userCollection = append(u.userCollection, *user)
	return user, nil
}

func (u *UserService) GetUser(id *int) (*models.User, error) {
	for _, user := range u.userCollection {
		if user.Id == *id {
			return &user, nil
		}
	}
	return nil, nil
}

func (u *UserService) UpdateUser(id *int, user *models.User) (*models.User, error) {
	for i, existingUser := range u.userCollection {
		if existingUser.Id == *id {
			u.userCollection[i] = *user
			return user, nil
		}
	}
	return nil, nil
}

func (u *UserService) GetAllUsers() *[]models.User {
	return &u.userCollection
}

func (u *UserService) DeleteUser(id *int) error {
	for i, existingUser := range u.userCollection {
		if existingUser.Id == *id {
			u.userCollection = append(u.userCollection[0:i], u.userCollection[i+1:]...)
		}
	}
	return nil
}
