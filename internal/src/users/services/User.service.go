// /internal/src/users/services/User.service.go
package services

import (
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	"go-rest/internal/src/users/ports/profile"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo       ports.UserRepository
	profileService profile.ProfileToUserAdapter
}

func NewUserService(userRepo ports.UserRepository, profileService profile.ProfileToUserAdapter) *UserService {
	return &UserService{
		UserRepo:       userRepo,
		profileService: profileService,
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
func (u *UserService) GetUserByID(id string) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(id)
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
func (u *UserService) DeleteUserByID(id string) error {
	err := u.UserRepo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserService) UpdateUserByID(id string, user domain.UserUpdateDTO) (domain.UserResponseDTO, error) {
	_, err := u.GetUserByID(id)
	if err != nil {
		return domain.UserResponseDTO{}, err
	}
	updateUser, err := u.UserRepo.UpdateUserByID(id, user)
	if err != nil {
		return domain.UserResponseDTO{}, err
	}
	return updateUser, nil
}
func (u *UserService) FormateUser(user domain.User) domain.UserResponseDTO {
	return domain.UserResponseDTO{
		Name:        user.Name,
		Email:       user.Email,
		Profile:     user.Profile,
		Preferences: user.Preferences,
	}
}
func (u *UserService) FormateUsers(users []domain.User) []domain.UserResponseDTO {
	var response []domain.UserResponseDTO
	for _, user := range users {
		response = append(response, domain.UserResponseDTO{
			Name:        user.Name,
			Email:       user.Email,
			Profile:     user.Profile,
			Preferences: user.Preferences,
		})
	}
	return response
}
