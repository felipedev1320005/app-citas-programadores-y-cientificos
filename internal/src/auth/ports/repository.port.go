// /internal/src/auth/ports/repository.port.go
package ports

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	UserEntity "go-rest/internal/src/users/domain"
)

// AuthRepository define las operaciones de repositorio necesarias para la autenticación.
// En este momento solo tiene un método para registrar usuarios.
type AuthRepository interface {
	Register(user dtos.AuthRegisterDOT) (UserEntity.User, error)
	Login(user dtos.AuthLoginDOT) (UserEntity.User, error) // Puedes agregar este método cuando implementes login
}
