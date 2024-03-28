package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runes/tools/envdata"
)

/*
SaveGame connects to DB to insert gameID and gameChampion (ChampionID) on it if it's not set already as evaluated in checkGameIsSet
That func is has a similar comparison to setTable, but they are not the same, each one using different tables
*/
func SaveGame(context *gin.Context, gameID string, championID int) {
	db := Connect()
	var table string
	if context.Request.URL.Path == "/play/guess/"+gameID {
		table = envdata.Env.TBGuess
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID {
		table = envdata.Env.TBMayhem
	}
	if checkGameIsSet(db, gameID, table) {
		return
	} else {
		_, errr := db.Exec(fmt.Sprintf("INSERT INTO `%s` (`gameID`, `ChampionID`) VALUES (?, ?);", table), gameID, championID)
		if errr != nil {
			log.Fatal("SaveGame() -> Error while saving gameID on DB")
		}
	}
	defer db.Close()
}

func SaveHangman(gameID, table string) {
	db := Connect()
	if checkGameIsSet(db, gameID, table) {
		return
	} else {
		_, err := db.Exec(fmt.Sprintf("INSERT INTO `%s` (`gameID`) VALUES ('%s');", table, gameID))
		if err != nil {
			log.Fatal("SaveGameHangman() -> Error while saving gameID on DB \n", err)
		}
	}
	defer db.Close()
}

/*
checkGameIsSet runs a count on DB table to check is a game is not already placed on it.
If found > 0 means that a game is already set, so we must get over this
*/
func checkGameIsSet(db *sql.DB, gameID, table string) bool {
	var found int
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM `%s` WHERE `gameID` = ?;", table), gameID).Scan(&found)
	if err != nil {
		log.Fatal("checkGameIsSet() -> Error while checking if game is already set")
	}
	return found > 0
}
