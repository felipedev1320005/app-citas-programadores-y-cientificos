// /internal/src/users/ports/UserHandler.port.go
package ports

import (
	"net/http"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	// GetUser(w http.ResponseWriter, r *http.Request)
	// UpdateUser(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)
}
