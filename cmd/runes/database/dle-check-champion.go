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

/*
 */
func CheckChampionNameByID(draw string) int {
	db := Connect()
	var drawChampionID int
	err := db.QueryRow("SELECT `ID` FROM `lol_Champions` WHERE NAME LIKE ?;", draw).Scan(&drawChampionID)
	if err != nil {
		log.Fatal("CheckChampionNameByID() -> error while checking ID for the given champion name")
	}
	defer db.Close()
	return drawChampionID
}
