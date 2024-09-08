package routes

import (
	handlers "go-rest/internal/src/users/handdlers"
	"go-rest/internal/src/users/repo"
	"go-rest/internal/src/users/services"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar repositorio, servicio y handler
	UserRepo := repo.UserRepository{}
	UserServices := services.UserService{UserRepo: &UserRepo}
	UserHandler := handlers.NewUserHandler(&UserServices)

	// Definir rutas
	r.HandleFunc("/users", UserHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.DeleteUserByID).Methods("DELETE")
	r.HandleFunc("/users/{id}", UserHandler.UpdateUserByID).Methods("PUT")

	return r
}
