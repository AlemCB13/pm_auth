package main

import (
	"log"
	"net/http"
	"pm_auth/src/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Configurar el enrutador
	r := mux.NewRouter()

	// Definir rutas
	apirRouter := r.PathPrefix("/api").Subrouter()
	apirRouter.HandleFunc("/register", handlers.Register).Methods("POST")
	apirRouter.HandleFunc("/login", handlers.Login).Methods("POST")
	apirRouter.HandleFunc("/logout", handlers.Logout).Methods("POST")
	apirRouter.HandleFunc("/2fa", handlers.TwoFactorAuth).Methods("POST")

	// Iniciar el servidor
	log.Println("Servidor de autenticación iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
