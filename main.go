package main

import (
    "log"
    "net/http"
    "lab6-web/database"
    "lab6-web/routes"
    "github.com/gorilla/handlers"
)

func main() {
    database.ConnectDB()

    r := routes.SetupRouter()

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", "https://miaplicacion.com"}), // URLs permitidas
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Métodos permitidos
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Cabeceras permitidas
    )

    log.Println("Servidor corriendo en el puerto 8100")

    log.Fatal(http.ListenAndServe(":8100", corsHandler(r)))
}