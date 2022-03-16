package main

import (
	"log"

	"github.com/anibalorellana/twittor/bd"
	"github.com/anibalorellana/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la Base de Datos :(")
		return
	}
	handlers.Manejadores()
}
