// /cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	AuthRoutes "go-rest/internal/src/auth/routes" // Corregido aquí
	repositoryconection "go-rest/internal/src/shared/repositoryConection"
	postgresRepo "go-rest/internal/src/shared/repositoryConection/posgress"
	UsersRoutes "go-rest/internal/src/users/routes" // Corregido aquí

	// Corregido aquí
	"github.com/gorilla/mux"
)

func main() {
	// Inicializa la conexión con la base de datos
	postgres := postgresRepo.PosgressRepositoryConection{}
	dbConection := repositoryconection.RepositoryConection{RepositoryConection: &postgres}
	dbConection.Conection()

	// Crear el router de Gorilla Mux
	r := mux.NewRouter()

	// Definir las rutas
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.PathPrefix("/auth").Handler(AuthRoutes.NewRouter())
	r.PathPrefix("/users").Handler(UsersRoutes.NewRouter())
	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Controladores
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la API con Gorilla Mux y Gorm!")
}
