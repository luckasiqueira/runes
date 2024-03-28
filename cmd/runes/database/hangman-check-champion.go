package database

import (
	"fmt"
	"log"
	"runes/tools/envdata"
	"strings"
)

func CheckHangmanChampion() []string {
	db := Connect()
	var championID int
	var c string
	table := envdata.Env.TBHangmanChampion
	err := db.QueryRow(fmt.Sprintf("SELECT `ChampionID` FROM `%s` WHERE 1;", table)).Scan(&championID)
	if err != nil {
		log.Fatal("CheckHangmanChampion() -> error when checking Hangman Champion.\n", err)
	}
	for i := range *ChampionsList {
		if championID == (*ChampionsList)[i].Champion.ID {
			c = (*ChampionsList)[i].Champion.Name
			break
		}
	}
	var champion []string
	for _, letter := range c {
		l := string(letter)
		l = strings.ToUpper(l)
		champion = append(champion, l)
	}
	defer db.Close()
	return champion
}
