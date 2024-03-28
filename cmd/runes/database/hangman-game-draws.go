package database

import (
	"fmt"
	"log"
	"runes/tools/envdata"
)

func CheckHangmanDraws(gameID string) []string {
	db := Connect()
	table := envdata.Env.TBHangmanDraws
	rows, err := db.Query(fmt.Sprintf("SELECT `draw` FROM `%s` WHERE `gameID` = '%s';", table, gameID))
	if err != nil {
		log.Fatal("CheckHangmanDraws() -> error while checking Hangman draws")
	}
	var Draws []string
	for rows.Next() {
		var draw string
		rows.Scan(&draw)
		Draws = append(Draws, draw)
	}
	defer db.Close()
	return Draws
}

func SaveHangmanGameDraws(gameID, draw string) {
	db := Connect()
	table := envdata.Env.TBHangmanDraws
	query := fmt.Sprintf("INSERT INTO `%s` (`id`, `draw`, `gameID`) VALUES (NULL, '%s', '%s');", table, draw, gameID)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("SaveHangmanGameDraws() -> error while saving Hangman draw\n", err)
	}
	defer db.Close()
}
