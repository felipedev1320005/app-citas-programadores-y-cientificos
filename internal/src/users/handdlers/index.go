// /internal/src/users/handlers/index.go
package handlers

import (
	"encoding/json"
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	UserService ports.UserService
	Validator   *validator.Validate
}

// NewUserHandler es el constructor que inicializa UserHandler con el validador.
func NewUserHandler(userService ports.UserService) *userHandler {
	return &userHandler{
		UserService: userService,
		Validator:   validator.New(), // Inicializa el validador aquí
	}
}

func (u *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userBody domain.UserCreateDTO

	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid JSON payload")
		return
	}

	err = u.Validator.Struct(userBody) // Usa el validador para validar la estructura
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	newUser, err := u.UserService.CreateUser(userBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
