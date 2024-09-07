// /cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	AuthHandlers "go-rest/internal/src/auth/handdlers" // Corregido aquí
	AuthRepo "go-rest/internal/src/auth/repo"          // Corregido aquí
	AuthServices "go-rest/internal/src/auth/services"  // Corregido aquí
	repositoryconection "go-rest/internal/src/shared/repositoryConection"
	postgresRepo "go-rest/internal/src/shared/repositoryConection/posgress"
	handlers "go-rest/internal/src/users/handdlers" // Corregido aquí
	"go-rest/internal/src/users/repo"
	"go-rest/internal/src/users/services"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa la conexión con la base de datos
	postgres := postgresRepo.PosgressRepositoryConection{}
	dbConection := repositoryconection.RepositoryConection{RepositoryConection: &postgres}
	dbConection.Conection()

	// Crear el router de Gorilla Mux
	r := mux.NewRouter()

	// Inicializar servicios, repositorios y handlers
	UserRepo := repo.UserRepository{}
	UserServices := services.UserService{UserRepo: &UserRepo}
	UserHandler := handlers.NewUserHandler(&UserServices) // Utiliza el constructor para crear el UserHandler

	// Inicializar servicios, repositorios y handlers de autenticación
	AuthRepo := AuthRepo.AuthRepository{}                        // Crear una instancia de AuthRepository
	AuthService := AuthServices.AuthService{AuthRepo: &AuthRepo} // Crear una instancia de AuthService
	AuthHandler := AuthHandlers.NewAuthHandler(&AuthService)     // Crear una instancia de AuthHandler
	// Definir las rutas
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users", UserHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", UserHandler.CreateUser).Methods("POST")
	//  register
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods("POST") // Agregar la ruta de registro de usuario
	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Controladores
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la API con Gorilla Mux y Gorm!")
}
