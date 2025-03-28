package models

import "gorm.io/gorm"

type Match struct {
    gorm.Model
	ID        int    `json:"id"`
    HomeTeam  string `json:"home_team"`
    AwayTeam  string `json:"away_team"`
    Score     string `json:"score"`
    matchDate      string `json:"date"`     
    Goals     int    `json:"goals"`      
    YellowCards   int    `json:"yellow_cards"`    
    RedCards      int    `json:"red_cards"`       
    ExtraTime     bool   `json:"extra_time"`
}