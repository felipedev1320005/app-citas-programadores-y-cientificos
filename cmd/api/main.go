package main

import (
	"fmt"
	"log"
	"net/http"

	repositoryconection "go-rest/internal/src/shared/repositoryConection"
	PosgresRepo "go-rest/internal/src/shared/repositoryConection/posgress"
	"go-rest/internal/src/users/handdlers"
	"go-rest/internal/src/users/repo"
	"go-rest/internal/src/users/services"

	"github.com/gorilla/mux"
)

func main() {
	postgres := PosgresRepo.PosgressRepositoryConection{}
	dbConection := repositoryconection.RepositoryConection{RepositoryConection: &postgres}
	dbConection.Conection()
	// Crear el router de Gorilla Mux
	r := mux.NewRouter()
	UserServices := services.UserService{}
	UserRepo := repo.UserRepository{UserService: &UserServices}
	UserHandler := handdlers.UserHandler{UserRepository: &UserRepo}
	// Definir las rutas
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users", UserHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", UserHandler.CreateUser).Methods("POST")

	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Controladores
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la API con Gorilla Mux y Gorm!")
}
