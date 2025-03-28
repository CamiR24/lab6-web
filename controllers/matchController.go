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
    rows, err := database.DB.Query(
        "SELECT id, home_team, away_team, score, date, goals, yellow_cards, red_cards, extra_time FROM matches")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var matches []models.Match

    for rows.Next() {
        var match models.Match
        if err := rows.Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.Score, &match.matchDate,
            &match.Goals, &match.YellowCards, &match.RedCards, &match.ExtraTime); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        matches = append(matches, match)
    }

    json.NewEncoder(w).Encode(matches)
}

func GetMatchByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var match models.Match
    err := database.DB.QueryRow(
        "SELECT id, home_team, away_team, score, date, goals, yellow_cards, red_cards, extra_time FROM matches WHERE id=$1", id).
        Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.Score, &match.matchDate,
            &match.Goals, &match.YellowCards, &match.RedCards, &match.ExtraTime)

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
        "INSERT INTO matches (home_team, away_team, score, date, goals, yellow_cards, red_cards, extra_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
        match.HomeTeam, match.AwayTeam, match.Score, match.matchDate,
        match.Goals, match.YellowCards, match.RedCards, match.ExtraTime,
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
        "UPDATE matches SET home_team=$1, away_team=$2, score=$3, date=$4, goals=$5, yellow_cards=$6, red_cards=$7, extra_time=$8 WHERE id=$9",
        match.HomeTeam, match.AwayTeam, match.Score, match.matchDate,
        match.Goals, match.YellowCards, match.RedCards, match.ExtraTime, id,
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
