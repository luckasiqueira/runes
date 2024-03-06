package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
SaveGame connects to DB to insert gameID and gameChampion (ChampionID) on it if it's not set already as evaluated in checkGameIsSet
*/
func SaveGame(context *gin.Context, gameID string) {
	db := Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("SaveGame() -> Error while loading .env file")
	}
	var table string
	if context.Request.URL.Path == "/play/guess/"+gameID {
		table = os.Getenv("TB_GUESS")
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID {
		table = os.Getenv("TB_MAYHEM")
	}
	if checkGameIsSet(db, gameID, table) {
		return
	} else {
		_, errr := db.Exec(fmt.Sprintf("INSERT INTO %s (`gameID`, `ChampionID`) VALUES (?, '55');", table), gameID)
		if errr != nil {
			log.Fatal("SaveGame() -> Error while saving gameID on DB")
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
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE `gameID` = ?;", table), gameID).Scan(&found)
	if err != nil {
		log.Fatal("checkGameIsSet() -> Error while checking if game is already set")
	}
	return found > 0
}
