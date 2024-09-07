// /internal/src/users/services/User.service.go
package services

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	"go-rest/internal/src/users/repo"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo ports.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepository(),
	}
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
func (u *UserService) GetUserByEmail(email string) (domain.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u *UserService) ComparePasswords(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
