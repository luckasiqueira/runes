package database

import (
	"fmt"
	"log"
)

func CheckGameChampion(gameID, table string) int {
	db := Connect()
	var championID int
	err := db.QueryRow(fmt.Sprintf("SELECT `ChampionID` FROM %s WHERE `gameID` = ?;", table), gameID).Scan(&championID)
	if err != nil {
		log.Fatal("CheckGameChampion() -> error while checking champion set for this game")
	}
	defer db.Close()
	return championID
}
