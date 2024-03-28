package database

import (
	"fmt"
	"log"
	"runes/tools/envdata"
)

func SaveHangmanChampion(championID int) {
	db := Connect()
	table := envdata.Env.TBHangmanChampion
	_, err := db.Exec(fmt.Sprintf("UPDATE `%s` SET `ID`='1',`ChampionID`=?;", table), championID)
	if err != nil {
		log.Fatal("SaveHangmanChampion() -> error while saving hangman champion to DB")
	}
	defer db.Close()
}
