package adapters

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
)

type authAdapter struct {
	userService ports.UserService
}

func NewAuthAdapter(userService ports.UserService) *authAdapter {
	return &authAdapter{userService: userService}
}

func (a *authAdapter) CreateUser(user domain.UserCreateDTO) (domain.User, error) {
	newUser, err := a.userService.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}
func (a *authAdapter) GetUserByEmail(email string) (domain.User, error) {
	user, err := a.userService.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (a *authAdapter) ComparePasswords(hashedPassword string, password string) error {
	err := a.userService.ComparePasswords(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}
