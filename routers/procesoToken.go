package routers

import (
	"errors"
	"strings"

	"github.com/anibalorellana/twittor/bd"
	"github.com/anibalorellana/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* Email valor del Email usado en todos los EndPoints */
var Email string

/* IDUsuario es el ID devueto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/* ProcesoToken  proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("anibalore")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Inválido")
	}
	return claims, false, string(""), err
}
