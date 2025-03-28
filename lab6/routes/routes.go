package routes

import (
    "github.com/gorilla/mux"
    "lab6/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/api/matches", controllers.GetAllMatches).Methods("GET")
    r.HandleFunc("/api/matches/{id}", controllers.GetMatchByID).Methods("GET")
    r.HandleFunc("/api/matches", controllers.CreateMatch).Methods("POST")
    r.HandleFunc("/api/matches/{id}", controllers.UpdateMatch).Methods("PUT")
    r.HandleFunc("/api/matches/{id}", controllers.DeleteMatch).Methods("DELETE")

    return r
}