package routers

import (
	"encoding/json"
	"net/http"

	"github.com/anibalorellana/twittor/bd"
	"github.com/anibalorellana/twittor/models"
)

/* Registro es la funcion para crear en la BD el registro de usuario - 17/03/2022 */
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	/*Valida correo requerido */
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	/*Valida que tenga contraseña de al menos 6 caracteres */
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}
	/* Valida que el usuario no exista con el correo*/
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese correo electrónico", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realiza el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
