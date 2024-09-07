// /internal/src/auth/handdlers/Auth.handdler.go
package handdlers

import (
	"encoding/json"
	dtos "go-rest/internal/src/auth/domain/DTOS"
	"go-rest/internal/src/auth/ports"
	"net/http"

	UserEntity "go-rest/internal/src/users/domain"

	"github.com/go-playground/validator/v10" // Paquete para validación de datos
)

type authHandler struct {
	AuthService ports.AuthService   // Servicio para manejar la lógica de autenticación
	Validator   *validator.Validate // Validador para verificar los datos de entrada
	UserService ports.UserService   // Adaptador para el servicio de usuarios
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
func NewAuthHandler(authService ports.AuthService, userService ports.UserService) (*authHandler, error) {
	return &authHandler{
		AuthService: authService,
		Validator:   validator.New(), // Inicializa el validador
		UserService: userService,
	}, nil
}

// Register maneja las solicitudes de registro de nuevos usuarios.
func (a *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userBody dtos.AuthRegisterDOT

	// Decodificar el cuerpo de la solicitud JSON en la estructura AuthRegisterDOT
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validar los datos del usuario
	err = a.Validator.Struct(userBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser := UserEntity.UserCreateDTO{
		Name:     userBody.Name,
		Email:    userBody.Email,
		Password: userBody.Password,
	}
	UserCreate, err := a.UserService.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	token, err := a.AuthService.GenerateTokenFromUser(UserCreate)
	// Registrar el nuevo usuario y generar un JWT
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Enviar el JWT como respuesta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Login maneja las solicitudes de inicio de sesión de los usuarios.
func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userBody dtos.AuthLoginDOT

	// Decodificar el cuerpo de la solicitud JSON en la estructura AuthLoginDOT
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid JSON payload")
		return
	}

	// Validar los datos del usuario
	err = a.Validator.Struct(userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user, err := a.UserService.GetUserByEmail(userBody.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}
	err = a.UserService.ComparePasswords(user.Password, userBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}
	token, err := a.AuthService.GenerateTokenFromUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token}) // Enviar el JWT como respuesta
}
