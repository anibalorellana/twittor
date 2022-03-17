package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword encripta password */
func EncriptarPassword(pass string) (string, error) {
	/* costo lo eleva al cuadrado, mayor el costo, mayor cantidad de pasada, mayor demora */
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
