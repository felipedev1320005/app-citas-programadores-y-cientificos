package services

import (
	UserEntity "go-rest/internal/src/users/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecret = "your-secret-key" // Cambia esto por una clave secreta segura

type AuthService struct {
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
func (a *AuthService) GenerateTokenFromUser(user UserEntity.User) (string, error) {
	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (a *AuthService) GetUserIDFromToken(tokenString string) (string, error) {
	token, err := a.VerifyToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	userID := claims["id"].(string)
	return userID, nil
}
