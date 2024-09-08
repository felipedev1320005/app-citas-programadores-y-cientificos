package routes

import (
	Adatapters "go-rest/internal/src/users/adapters"
	handlers "go-rest/internal/src/users/handdlers"
	"go-rest/internal/src/users/repo"
	ProfileRepo "go-rest/internal/src/users/repo/profile"
	"go-rest/internal/src/users/services"
	ProfileService "go-rest/internal/src/users/services/profile"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar repositorio, servicio y handler
	UserRepo := repo.UserRepository{}
	UserServices := services.NewUserService(&UserRepo, Adatapters.NewProfileToUserAdapter(ProfileService.NewProfileService(ProfileRepo.NewProfileRepository())))
	ProfileRepo := ProfileRepo.NewProfileRepository()
	ProfileSerivce := ProfileService.NewProfileService(ProfileRepo)
	UserHandler := handlers.NewUserHandler(UserServices, ProfileSerivce)
	// Definir rutas
	r.HandleFunc("/users", UserHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler.DeleteUserByID).Methods("DELETE")
	r.HandleFunc("/users/{id}", UserHandler.UpdateUserByID).Methods("PUT")

	return r
}
