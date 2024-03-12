package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setTable(playingMode string) string {
	errr := godotenv.Load()
	if errr != nil {
		log.Fatal("setTable() -> error while loading .env file")
	}
	var table string
	if playingMode == "Guess" {
		table = os.Getenv("TB_GUESS_DRAWS")
	} else if playingMode == "Mayhem" {
		table = os.Getenv("TB_MAYHEM_DRAWS")
	}
	return table
}

/*
SaveDraw connects to DB and place the given championID on it, in order to save all shots for each game
*/
func SaveDraw(gameID, playingMode string, championID int) {
	table := setTable(playingMode)
	db := Connect()
	_, err := db.Exec(fmt.Sprintf("INSERT INTO `%s` (`id`, `draw`, `gameID`) VALUES (NULL, '%d', '%s');", table, championID, gameID))
	if err != nil {
		log.Fatal("SaveDraw() -> error while saving draw on BD")
	}
	defer db.Close()
}

func CheckDraws() {

}
