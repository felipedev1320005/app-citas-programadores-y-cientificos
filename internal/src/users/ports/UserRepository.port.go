// /internal/src/users/ports/UserRepository.port.go
package ports

import "go-rest/internal/src/users/domain"

type UserRepository interface {
	CreateUser(user domain.UserCreateDTO) (domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	DeleteUserByID(id string) error
	GetUserByID(id string) (domain.User, error)
	UpdateUserByID(id string, user domain.UserUpdateDTO) (domain.UserResponseDTO, error)
}
type ProfileRepository interface {
	CreateProfile(profile domain.ProfileCreateDTO) (domain.ProfileResponseDTO, error)
}
