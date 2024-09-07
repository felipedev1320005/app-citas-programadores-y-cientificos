// /internal/src/shared/middleware/AuthMiddleware.go
package middleware

import (
	"go-rest/internal/src/auth/ports" // Ajusta el import a la ubicación real de tu puerto AuthService
	"net/http"
	"strings"
)

// JWTMiddleware es un middleware para autenticar las solicitudes mediante JWT
func JWTMiddleware(next http.Handler, authService ports.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		// Asegurarse de que el token comience con "Bearer "
		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extraer el token
		tokenString = tokenString[len("Bearer "):]

		// Verificar el token
		_, err := authService.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Continuar con la siguiente función handler
		next.ServeHTTP(w, r)
	})
}
