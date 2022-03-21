package models

/*RespuestaLogin tiene el token que se devuelve con el inicio de sesión */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
