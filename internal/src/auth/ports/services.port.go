// /internal/src/auth/port/services.port.go
package ports

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	UserEntity "go-rest/internal/src/users/domain"

	"github.com/dgrijalva/jwt-go"
)

// AuthService define los métodos necesarios para la lógica de negocio relacionada con la autenticación.
// En este momento solo tiene un método para registrar usuarios.
type AuthService interface {
	Register(user dtos.AuthRegisterDOT) (string, error)
	Login(user dtos.AuthLoginDOT) (string, error) // Puedes agregar este método cuando implementes login
	VerifyToken(tokenString string) (*jwt.Token, error)
	GenerateTokenFromUser(user UserEntity.User) (string, error)
	GetUserIDFromToken(tokenString string) (string, error)
}
