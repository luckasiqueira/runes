package dle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"runes/cmd/runes/database"
)

var c = database.CheckDailyChampion()
var dailyChampion *database.ChampionLOL = &c

/*
PlayDLE is where game starts, implementing some initial evaluations
gameID collects the gameID param from the URL, to identify from the URL what game mode is user playing
dailyDraw is executed if game mode is "guess" and user must guess champion set on pointer dailyChampion
*/
func PlayDLE(context *gin.Context, draw string) {
	gameID := context.Param("gameID")
	var championID int
	var champion database.ChampionLOL
	table := database.SetTable(context)
	var playingMode string
	if context.Request.URL.Path == "/try/guess/"+gameID {
		championID = (*dailyChampion).ID
		playingMode = "Guess"
	} else if context.Request.URL.Path == "/try/mayhem/"+gameID {
		DraftDailyChampion()
		championID = database.CheckGameChampion(gameID, table)
		playingMode = "Mayhem"
	}
	drawChampion := database.CheckChampionDrawed(draw)
	compare(championID, drawChampion, champion)
	go database.SaveDraw(gameID, playingMode, drawChampion.ID)
}

/*
compare compares all characteristics for the given champion drawed with the defined champion for this game.
If a characteristic is correct, it will set a Status as true, which will be used to indicate to player if that shot is correct, partially correct or wrong.
*/
func compare(championID int, drawChampion, champion database.ChampionLOL) {
	if drawChampion.ID == championID {
		fmt.Println("WIN")
	} else {
		if drawChampion.Gender == champion.Gender {
			//
		}
	}
}

/*
dailyDraw runs every 00:00h (here, set as 12), when it Draws a new champion from DB and saves onto dailyChampion pointer
This pointer will be used to compare user shots fast, since no DB comparision will be needed
*/
func DraftDailyChampion() {
	cronCycle := "0 0 * * *"
	job := cron.New()
	job.AddFunc(cronCycle, func() {
		c = database.DrawChampion()
		database.SaveDailyChampion(c.ID)
		for i := range *database.ChampionsList {
			if c.ID == (*database.ChampionsList)[i].Champion.ID {
				dailyChampion = &c
				break
			}
		}
	})
	job.Start()
	job.Run()
}
