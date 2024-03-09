package database

import "log"

/*
Struct for each champion
*/
type ChampionLOL struct {
	ID       int
	Avatar   string
	Name     string
	Gender   string
	Role     string
	Race     string
	Resource string
	Range    string
	Region   string
	Release  string
}

/*
Status for each guess where "found" is correct and "partial" is partially correct
*/
type DrawStatus struct {
	GenderFound   bool
	RoleFound     bool
	RolePartial   bool
	RaceFound     bool
	RacePartial   bool
	ResourceFound bool
	RangeFound    bool
	RangePartial  bool
	RegionFound   bool
	RegionPartial bool
	ReleaseFound  bool
	ReleaseUp     bool
	ReleaseDown   bool
}

/*
Includes all user guesses, settiing all guessed champions, status for each guess and if games has been won
*/
type Draws struct {
	Champion ChampionLOL
	Status   DrawStatus
	Won      bool
}

/*
Saving in a struct pointer all champions from DB
*/
var c = ListChampions()
var ChampionsList *[]Draws = &c

/*
ListChampions connects to DB and get all champion's info, saving it onto ChampionsList
This approach can reduce new DB connections, and improve comparison speed
*/
func ListChampions() []Draws {
	db := Connect()
	rows, err := db.Query("SELECT * FROM`lol_Champions`")
	if err != nil {
		log.Fatal("ListChampions() -> error while getting all champions info from DB")
	}
	var championsList []Draws
	for rows.Next() {
		eachChampion := Draws{}
		rows.Scan(
			&eachChampion.Champion.ID,
			&eachChampion.Champion.Name,
			&eachChampion.Champion.Gender,
			&eachChampion.Champion.Role,
			&eachChampion.Champion.Race,
			&eachChampion.Champion.Resource,
			&eachChampion.Champion.Range,
			&eachChampion.Champion.Region,
			&eachChampion.Champion.Release,
			&eachChampion.Champion.Avatar,
		)
		championsList = append(championsList, eachChampion)
	}
	defer db.Close()
	return championsList
}
