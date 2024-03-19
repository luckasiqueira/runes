package database

import (
	"fmt"
	"log"
	"runes/tools/envdata"
)

/*
saveDailyChampion connects to DB and updates the championID with the new one generated today
*/
func SaveDailyChampion(championID int) {
	db := Connect()
	table := envdata.Env.TBGuessChampion
	_, err := db.Exec(fmt.Sprintf("UPDATE `%s` SET `ID`='1',`ChampionID`=?;", table), championID)
	if err != nil {
		log.Fatal("saveDailyChampion() -> error while saving daily champion to DB")
	}
	defer db.Close()
}
