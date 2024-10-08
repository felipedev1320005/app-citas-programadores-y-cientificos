// /internal/src/users/handlers/index.go
package handlers

import (
	"encoding/json"
	"go-rest/internal/src/users/domain"
	"go-rest/internal/src/users/ports"
	PortProfile "go-rest/internal/src/users/ports/profile"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type userHandler struct {
	UserService    ports.UserService
	Validator      *validator.Validate
	profileService PortProfile.ProfileService
}

// NewUserHandler es el constructor que inicializa UserHandler con el validador.
func NewUserHandler(userService ports.UserService, profileService PortProfile.ProfileService) *userHandler {
	return &userHandler{
		UserService: userService,
		Validator:   validator.New(), // Inicializa el validador aquí
	}
}
func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u.UserService.FormateUsers(users))
}
func (u *userHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Obtiene el id de la URL
	id := mux.Vars(r)["id"]

	err := u.UserService.DeleteUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted successfully")
}

func (u *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := u.UserService.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u.UserService.FormateUser(user))
}
func (u *userHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userUpdate := domain.UserUpdateDTO{}
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = u.Validator.Struct(userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user, err := u.UserService.UpdateUserByID(id, userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
