package models

import (
    "gorm.io/gorm"
)

type Match struct {
    gorm.Model
	ID        int    `json:"id"`
    HomeTeam  string `json:"homeTeam"`
    AwayTeam  string `json:"awayTeam"`
    MatchDate   string `json:"matchDate"`   
    Goals     int    `json:"goals"`      
    YellowCards   int    `json:"yellow_cards"`    
    RedCards      int    `json:"red_cards"`       
    ExtraTime     bool   `json:"extra_time"`
}