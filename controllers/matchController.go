package controllers

import (
    "encoding/json"  
    "net/http"       
    "strconv"        
    "github.com/gorilla/mux" 
    "lab6/database"
    "lab6/models"
)


func GetAllMatches(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, home_team, away_team, date FROM matches")
    if err != nil {
        http.Error(w, "Error al obtener los partidos: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var matches []models.Match

    for rows.Next() {
        var match models.Match
        if err := rows.Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.MatchDate); err != nil {
            http.Error(w, "Error al escanear los partidos: "+err.Error(), http.StatusInternalServerError)
            return
        }
        matches = append(matches, match)
    }

    if len(matches) == 0 {
        w.WriteHeader(http.StatusOK) // 200 OK
        json.NewEncoder(w).Encode([]models.Match{})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    
    json.NewEncoder(w).Encode(matches)
}


func GetMatchByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var match models.Match
    err := database.DB.QueryRow(
        "SELECT id, home_team, away_team, date FROM matches WHERE id=$1", id).
        Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.MatchDate)

    if err != nil {
        http.Error(w, "Partido no encontrado", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(match)
}

func CreateMatch(w http.ResponseWriter, r *http.Request) {
    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err := database.DB.Exec(
        "INSERT INTO matches (home_team, away_team, date) VALUES ($1, $2, $3)",
        match.HomeTeam, match.AwayTeam, match.MatchDate,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


func UpdateMatch(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err := database.DB.Exec(
        "UPDATE matches SET home_team=$1, away_team=$2, date=$3 WHERE id=$4",
        match.HomeTeam, match.AwayTeam, match.MatchDate, id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}


func DeleteMatch(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"]) //convierte el id de string a int

    _, err := database.DB.Exec("DELETE FROM matches WHERE id=$1", id)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func AddGoal(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err = database.DB.Exec(
        "UPDATE matches SET goals = goals + 1 WHERE id = $1",
        id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)

}

func AddYellowCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err = database.DB.Exec(
        "UPDATE matches SET yellow_cards = yellow_cards + 1 WHERE id = $1",
        id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    
}

func AddRedCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err = database.DB.Exec(
        "UPDATE matches SET red_cards = red_cards + 1 WHERE id = $1",
        id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    
}

func AddExtraTime(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    var match models.Match
    json.NewDecoder(r.Body).Decode(&match)

    _, err = database.DB.Exec(
        "UPDATE matches SET extra_time = TRUE WHERE id = $1",
        id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    
}
