package middlew

import (
	"net/http"

	"github.com/anibalorellana/twittor/bd"
)

/* ChequeoBD es el middlew que me permite conocer el estado de la BD - 17/03/2022 */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión Perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
