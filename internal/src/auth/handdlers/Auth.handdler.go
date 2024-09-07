// /internal/src/auth/handdlers/Auth.handdler.go
package handdlers

import (
	"encoding/json"
	dtos "go-rest/internal/src/auth/domain/DTOS"
	"go-rest/internal/src/auth/ports"
	"net/http"

	"github.com/go-playground/validator/v10" // Paquete para validación de datos
)

type AuthHandler struct {
	AuthService ports.AuthService   // Servicio para manejar la lógica de autenticación
	Validator   *validator.Validate // Validador para verificar los datos de entrada
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
func NewAuthHandler(authService ports.AuthService) (*AuthHandler, error) {
	return &AuthHandler{
		AuthService: authService,
		Validator:   validator.New(), // Inicializa el validador
	}, nil
}

// Register maneja las solicitudes de registro de nuevos usuarios.
func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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

	// Registrar el nuevo usuario y generar un JWT
	token, err := a.AuthService.Register(userBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Enviar el JWT como respuesta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Login maneja las solicitudes de inicio de sesión de los usuarios.
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
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

	// Iniciar sesión del usuario y generar un JWT
	token, err := a.AuthService.Login(userBody)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token}) // Enviar el JWT como respuesta
}
