package routes

import (
	"go-rest/internal/src/auth/handdlers"
	"go-rest/internal/src/auth/services"

	UserAdapter "go-rest/internal/src/users/adapters"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar repositorio, servicio y handler
	AuthService := services.AuthService{}
	UserAdapter := UserAdapter.NewAuthAdapter()
	AuthHandler, errhand := handdlers.NewAuthHandler(&AuthService, UserAdapter)
	if errhand != nil {
		panic(errhand)
	}
	// Definir rutas
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", AuthHandler.Login).Methods("POST")

	return r
}
