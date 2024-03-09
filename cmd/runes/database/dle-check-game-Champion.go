package database

import (
	"fmt"
	"log"
)

/*
CheckGameChampion looks on DB to verify what champion ID is set for the  given gameID and returns a INT
*/
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
