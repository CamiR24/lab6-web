package main

import (
    "log"
    "net/http"
    "lab6/database"
    "lab6/routes"
)

func main() {
	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", "https://miaplicacion.com"}), // URLs permitidas
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Métodos permitidos
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Cabeceras permitidas
    )

    database.ConnectDB()
    r := routes.SetupRouter()

    log.Println("Servidor corriendo en el puerto 8100")
    log.Fatal(http.ListenAndServe(":8100", r))
}