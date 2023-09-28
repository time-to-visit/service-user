package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// EncriptarPassword encripta una contraseña utilizando bcrypt.
func EncriptarPassword(password string) (string, error) {
	// Generar un hash de la contraseña con un costo de trabajo (work factor) predeterminado
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerificarPassword compara una contraseña en texto claro con su versión encriptada.
func VerificarPassword(password, hashedPassword string) error {
	// Comparar la contraseña en texto claro con su versión encriptada
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
