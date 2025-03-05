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
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("POST")
	r.HandleFunc("/2fa", handlers.TwoFactorAuth).Methods("POST")
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	// Iniciar el servidor
	log.Println("Servidor de autenticación iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
