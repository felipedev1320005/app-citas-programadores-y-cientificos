package routes

import (
	Adatapters "go-rest/internal/src/users/adapters"
	handlers "go-rest/internal/src/users/handdlers"
	"go-rest/internal/src/users/repo"
	ProfileRepo "go-rest/internal/src/users/repo/profile"
	ProfileService "go-rest/internal/src/users/services/profile"
	UserService "go-rest/internal/src/users/services/users"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Repositorios
	UserRepo := repo.NewUserRepository()
	ProfileRepo := ProfileRepo.NewProfileRepository()
	// Servicios
	ProfileSerivce := ProfileService.NewProfileService(ProfileRepo)
	ProfileToUserAdapater := Adatapters.NewProfileToUserAdapter(ProfileSerivce) // Adapter para convertir el servicio de perfil en un servicio para usuario
	UserServices := UserService.NewUserService(UserRepo, ProfileToUserAdapater)
	UserHandler := handlers.NewUserHandler(UserServices, ProfileSerivce)
	// Definir rutas
	r.HandleFunc("/users", UserHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.DeleteUserByID).Methods("DELETE")
	r.HandleFunc("/users/{id}", UserHandler.UpdateUserByID).Methods("PUT")

	return r
}
