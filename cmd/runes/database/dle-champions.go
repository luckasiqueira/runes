package database

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
Salvando em um ponteiro de struct todos os campeões cadastrados no banco de dados
var c = ListChampions()
var ChampionsList *[]Draws = &c
*/
