package services

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	"go-rest/internal/src/auth/ports"
	UserEntity "go-rest/internal/src/users/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecret = "your-secret-key" // Cambia esto por una clave secreta segura

type AuthService struct {
	AuthRepo ports.AuthRepository
}

func (a *AuthService) Register(user dtos.AuthRegisterDOT) (string, error) {
	newUser, err := a.AuthRepo.Register(user)
	if err != nil {
		return "", err
	}

	token, err := generateToken(newUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (a *AuthService) Login(user dtos.AuthLoginDOT) (string, error) {
	existingUser, err := a.AuthRepo.Login(user)
	if err != nil {
		return "", err
	}

	token, err := generateToken(existingUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateToken(user UserEntity.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *AuthService) VerifyToken(tokenString string) (*jwt.Token, error) {
	var secretKey = "your_secret_key"

	// Parse el token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica el método de firma del token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid // Cambié esto para devolver un error específico
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
