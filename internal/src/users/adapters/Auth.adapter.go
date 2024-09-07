package adapters

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	"go-rest/internal/src/users/services"
)

type authAdapter struct {
	UserService ports.UserService
}

func NewAuthAdapter() *authAdapter {
	return &authAdapter{UserService: services.NewUserService()}
}

func (a *authAdapter) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser, err := a.UserService.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
func (a *authAdapter) GetUserByEmail(email string) (domain.User, error) {
	user, err := a.UserService.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
