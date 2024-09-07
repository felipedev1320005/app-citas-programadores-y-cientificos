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
	AuthService := services.AuthService{AuthRepo: &AuthRepo}
	AuthHandler := handdlers.AuthHandler{AuthService: &AuthService}

	// Definir rutas
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")

	return r
}
