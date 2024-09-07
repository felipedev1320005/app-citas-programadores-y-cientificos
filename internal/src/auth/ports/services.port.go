// /internal/src/auth/port/services.port.go
package ports

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	UserEntity "go-rest/internal/src/users/domain"
)

// AuthService define los métodos necesarios para la lógica de negocio relacionada con la autenticación.
// En este momento solo tiene un método para registrar usuarios.
type AuthService interface {
	Register(user dtos.AuthRegisterDOT) (UserEntity.User, error)
	// Login(user dtos.AuthLoginDOT) (UserEntity.User, error) // Puedes agregar este método cuando implementes login
}
