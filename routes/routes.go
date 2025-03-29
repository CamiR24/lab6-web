package routes

import (
    "github.com/gorilla/mux"
    "lab6/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    r.Use(controllers.EnableCORS) // Agregar middleware de CORS

    r.HandleFunc("/api/matches", controllers.GetAllMatches).Methods("GET")
    r.HandleFunc("/api/matches/{id}", controllers.GetMatchByID).Methods("GET")
    r.HandleFunc("/api/matches", controllers.CreateMatch).Methods("POST")
    r.HandleFunc("/api/matches/{id}", controllers.UpdateMatch).Methods("PUT")
    r.HandleFunc("/api/matches/{id}", controllers.DeleteMatch).Methods("DELETE")
    r.HandleFunc("/api/matches/{id}/goals", controllers.AddGoal).Methods("PATCH")
    r.HandleFunc("/api/matches/{id}/yellowcards", controllers.AddYellowCard).Methods("PATCH")
    r.HandleFunc("/api/matches/{id}/redcards", controllers.AddRedCard).Methods("PATCH")
    r.HandleFunc("/api/matches/{id}/extratime", controllers.AddExtraTime).Methods("PATCH")

    return r
}