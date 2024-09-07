// /internal/src/auth/domain/DTOS/Register.dto.go
package dtos

// AuthRegisterDOT define el formato para los datos de registro.
// Utiliza etiquetas `validate` para asegurar que los campos no estén vacíos y que el email tenga un formato válido.
type AuthRegisterDOT struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
