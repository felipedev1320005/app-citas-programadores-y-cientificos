// /internal/src/auth/services/Auth.service.go
package services

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	"go-rest/internal/src/auth/ports"
	UserEntity "go-rest/internal/src/users/domain"
)

type AuthService struct {
	UserRepo ports.AuthRepository // Repositorio para manejar operaciones de usuario
}

// Register usa el repositorio para registrar un nuevo usuario.
func (a *AuthService) Register(user dtos.AuthRegisterDOT) (UserEntity.User, error) {
	newUser, err := a.UserRepo.Register(user)
	if err != nil {
		return UserEntity.User{}, err
	}
	return newUser, nil
}
