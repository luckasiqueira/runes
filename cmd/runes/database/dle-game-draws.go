package database

import (
	"fmt"
	"log"
	"runes/tools/envdata"
)

/*
Set tables verify the current game mode to check what's Draw's related DB table to that mode
Once evaluation is done, a that related table name is returned, to be used on DB queries
*/
func setTable(playingMode string) string {
	var table string
	if playingMode == "guess" {
		table = envdata.Env.TBGuessDraws
	} else if playingMode == "mayhem" {
		table = envdata.Env.TBMayhemDraws
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
	rows, err := db.Query(fmt.Sprintf("SELECT `draw` FROM `%s` WHERE `gameID` = '%s' ORDER BY `id` DESC;", table, gameID))
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
