package database

import "log"

/*
saveDailyChampion connects to DB and updates the championID with the new one generated today
*/
func saveDailyChampion(championID int) {
	db := Connect()
	_, err := db.Exec("UPDATE `lol_Game_DailyChampion` SET `ID`='1',`ChampionID`=?;", championID)
	if err != nil {
		log.Fatal("saveDailyChampion() -> error while saving daily champion to DB")
	}
	defer db.Close()
}
