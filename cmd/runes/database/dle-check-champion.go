package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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
CheckChampionNameByID
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

/*
CheckDailyChampion will connect to DB in order to get the already defined championID
That function is only called once, to set the initial value for the dailyChampion pointer
*/
func CheckDailyChampion() ChampionLOL {
	db := Connect()
	var champion ChampionLOL
	var dailyChampionID int
	err := godotenv.Load()
	if err != nil {
		log.Fatal("CheckDailyChampion() -> error while loading .env file")
	}
	table := os.Getenv("TB_GUESS_CHAMPION")
	err = db.QueryRow(fmt.Sprintf("SELECT `championID` FROM %s WHERE 1;", table)).Scan(&dailyChampionID)
	if err != nil {
		log.Fatal("CheckDailyChampion() -> error while checking daily champion")
	}
	for i := range *ChampionsList {
		if dailyChampionID == (*ChampionsList)[i].Champion.ID {
			champion = (*ChampionsList)[i].Champion
		}
	}
	defer db.Close()
	return champion
}
