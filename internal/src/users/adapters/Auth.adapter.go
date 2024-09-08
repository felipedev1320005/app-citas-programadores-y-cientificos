package adapters

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	UserRepo "go-rest/internal/src/users/repo"
	ProfileRepo "go-rest/internal/src/users/repo/profile"
	"go-rest/internal/src/users/services"
	ProfileService "go-rest/internal/src/users/services/profile"
)

type authAdapter struct {
	userService ports.UserService
}

func NewAuthAdapter() *authAdapter {
	return &authAdapter{userService: services.NewUserService(UserRepo.NewUserRepository(), NewProfileToUserAdapter(ProfileService.NewProfileService(ProfileRepo.NewProfileRepository())))}
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
