package database

/*
Connect onto Database to select a random champiom and atributes it to championSelect,
which is a instance of ChampionLOL struct

func DrawChampion() ChampionLOL {
	db := Connect()
	var championSelected ChampionLOL
	err := db.QueryRow("SELECT * FROM `lol_Champions` ORDER BY RAND() LIMIT 1;").Scan(
		&championSelected.ID,
		&championSelected.Name,
		&championSelected.Gender,
		&championSelected.Role,
		&championSelected.Race,
		&championSelected.Resource,
		&championSelected.Range,
		&championSelected.Region,
		&championSelected.Release,
		&championSelected.Avatar)
	if err != nil {
		log.Fatal("DrawChampion -> error while performing db.QueryRow")
	}
	defer db.Close()
	return championSelected
}
*/
