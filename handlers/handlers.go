package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/anibalorellana/twittor/middlew"
	"github.com/anibalorellana/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo puerto, el handler y pongo a escuchar mi servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
