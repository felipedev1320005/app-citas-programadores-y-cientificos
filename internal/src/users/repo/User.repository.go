// /internal/src/users/repo/User.repository.go
package repo

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
)

type UserRepository struct {
	UserService ports.UserService
}

func (u *UserRepository) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser, err := u.UserService.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
func (u *UserRepository) GetUsers() ([]domain.User, error) {
	user, err := u.UserService.GetUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return user, nil
}
