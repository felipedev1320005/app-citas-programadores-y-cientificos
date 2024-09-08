// /internal/src/users/ports/UserRepository.port.go
package ports

import "go-rest/internal/src/users/domain"

type UserRepository interface {
	CreateUser(user domain.UserCreateDTO) (domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	DeleteUserByID(id string) error
	GetUserByID(id string) (domain.User, error)
	// GetUser(id string) (domain.User, error)
	// UpdateUser(id string, user domain.UserUpdateDTO) (domain.User, error)
	// DeleteUser(id string) error
}
