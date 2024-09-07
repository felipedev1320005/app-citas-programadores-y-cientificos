// /internal/src/auth/port/services.port.go
package ports

import (
	UserEntity "go-rest/internal/src/users/domain"

	"github.com/dgrijalva/jwt-go"
)

// AuthService define los métodos necesarios para la lógica de negocio relacionada con la autenticación.
// En este momento solo tiene un método para registrar usuarios.
type AuthService interface {
	VerifyToken(tokenString string) (*jwt.Token, error)
	GenerateTokenFromUser(user UserEntity.User) (string, error)
	GetUserIDFromToken(tokenString string) (string, error)
}
