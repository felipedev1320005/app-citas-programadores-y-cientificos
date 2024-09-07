package routes

import (
	"go-rest/internal/src/auth/handdlers"
	"go-rest/internal/src/auth/repo"
	"go-rest/internal/src/auth/services"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar repositorio, servicio y handler
	AuthRepo := repo.AuthRepository{}
	AuthService := services.AuthService{
		AuthRepo: &AuthRepo,
	}
	AuthHandler, errhand := handdlers.NewAuthHandler(&AuthService)
	if errhand != nil {
		panic(errhand)
	}
	// Definir rutas
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", AuthHandler.Login).Methods("POST")

	return r
}
