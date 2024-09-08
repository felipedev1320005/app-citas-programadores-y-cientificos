package routes

import (
	"go-rest/internal/src/auth/handdlers"
	"go-rest/internal/src/auth/services"

	UserAdapter "go-rest/internal/src/users/adapters"

	UserRepo "go-rest/internal/src/users/repo"
	ProfileRepo "go-rest/internal/src/users/repo/profile"
	ProfileService "go-rest/internal/src/users/services/profile"
	UserService "go-rest/internal/src/users/services/users"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar repositorio, servicio y handler
	AuthService := services.AuthService{}
	ProfileRepository := ProfileRepo.NewProfileRepository()
	ProfileService := ProfileService.NewProfileService(ProfileRepository)
	UserRepository := UserRepo.NewUserRepository()
	UserService := UserService.NewUserService(UserRepository, ProfileService)
	UserAdapter := UserAdapter.NewAuthAdapter(UserService)
	AuthHandler, errhand := handdlers.NewAuthHandler(&AuthService, UserAdapter)
	if errhand != nil {
		panic(errhand)
	}
	// Definir rutas
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", AuthHandler.Login).Methods("POST")

	return r
}
