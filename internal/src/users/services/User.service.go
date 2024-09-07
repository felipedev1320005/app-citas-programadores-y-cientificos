// /internal/src/users/services/User.service.go
package services

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
)

type UserService struct {
	UserRepo ports.UserRepository
}

func (u *UserService) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser, err := u.UserRepo.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
func (u *UserService) GetUsers() ([]domain.User, error) {
	user, err := u.UserRepo.GetUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return user, nil
}
