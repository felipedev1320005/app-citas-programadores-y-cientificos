package services

import (
	"go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/domain"
)

type UserService struct{}

func (u *UserService) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	createUser := posgress.Db.Create(&newUser)
	err := createUser.Error
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
func (u *UserService) GetUsers() ([]domain.User, error) {
	var users []domain.User
	getUsers := posgress.Db.Find(&users)
	err := getUsers.Error
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
