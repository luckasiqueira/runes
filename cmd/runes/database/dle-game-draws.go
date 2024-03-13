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
	if playingMode == "guess" {
		table = os.Getenv("TB_GUESS_DRAWS")
	} else if playingMode == "mayhem" {
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

/*
CheckDraws opens a db connection and looks for each draw for the given gameID
For each found draw, performs a loop over ChampionsList,
where an instance of gameDraw (from type Draw) is set for every found champion that corresponds to the found championID
*/
func CheckDraws(gameID, playingMode string) []Draws {
	var gameDraws []Draws
	table := setTable(playingMode)
	db := Connect()
	rows, err := db.Query(fmt.Sprintf("SELECT `draw` FROM `%s` WHERE `gameID` = '%s';", table, gameID))
	if err != nil {
		log.Fatal("CheckDraws -> error while list all gameDraws")
	}
	for rows.Next() {
		var gameDraw Draws
		var championID int
		rows.Scan(&championID)
		for i := range *ChampionsList {
			if championID == (*ChampionsList)[i].Champion.ID {
				gameDraw = (*ChampionsList)[i]
				gameDraws = append(gameDraws, gameDraw)
				break
			}
		}
	}
	defer db.Close()
	return gameDraws
}
