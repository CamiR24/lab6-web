package database

import (
    "database/sql"      
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    DB, err = sql.Open("postgres", "postgres://Camila:CR2343universidad@localhost:5432/Liga?sslmode=disable")

    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    }

    log.Println("Conexión exitosa a la base de datos")
}

