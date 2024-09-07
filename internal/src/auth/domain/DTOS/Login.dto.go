// /internal/src/auth/domain/DTOS/Login.dto.go
package dtos

// AuthLoginDOT define el formato para los datos de login.
// Utiliza etiquetas `validate` para asegurar que los campos no estén vacíos y que el email tenga un formato válido.
type AuthLoginDOT struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
