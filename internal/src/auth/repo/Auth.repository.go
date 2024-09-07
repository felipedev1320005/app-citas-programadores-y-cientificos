// /internal/src/auth/repo/Auth.repository.go
package repo

import (
	dtos "go-rest/internal/src/auth/domain/DTOS"
	"go-rest/internal/src/shared/repositoryConection/posgress"
	UserEntity "go-rest/internal/src/users/domain"

	"golang.org/x/crypto/bcrypt" // Paquete para encriptar contraseñas
)

type AuthRepository struct{}

// Register crea un nuevo usuario, encriptando la contraseña antes de guardarla en la base de datos.
func (a *AuthRepository) Register(user dtos.AuthRegisterDOT) (UserEntity.User, error) {
	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserEntity.User{}, err
	}

	newUser := UserEntity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword), // Guardar la contraseña encriptada
	}

	// Guardar el nuevo usuario en la base de datos
	createUser := posgress.Db.Create(&newUser)
	err = createUser.Error
	if err != nil {
		return UserEntity.User{}, err
	}
	return newUser, nil
}
func (a *AuthRepository) Login(user dtos.AuthLoginDOT) (UserEntity.User, error) {
	// Buscar el usuario por su correo electrónico
	var existingUser UserEntity.User
	getUser := posgress.Db.Where("email = ?", user.Email).First(&existingUser)
	err := getUser.Error
	if err != nil {
		return UserEntity.User{}, err
	}

	// Comparar la contraseña ingresada con la contraseña almacenada
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return UserEntity.User{}, err
	}

	return existingUser, nil
}
