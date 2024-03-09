package dle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"runes/cmd/runes/database"
)

var dailyChampion *database.ChampionLOL

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
	if context.Request.URL.Path == "/try/guess/"+gameID {
		championID = (*dailyChampion).ID
	} else if context.Request.URL.Path == "/try/mayhem/"+gameID {
		championID = database.CheckGameChampion(gameID, table)
	}
	fmt.Sprint(championID, champion)
	fmt.Println(draw)
}

/*
dailyDraw runs every 00:00 (here, set as 12), when it Draws a new champion from DB and saves onto dailyChampion pointer
This pointer will be used to compare user shots fast, since no DB comparision will be needed
*/
func DraftDailyChampion() {
	cronCycle := "0 0 * * *"
	job := cron.New()
	job.AddFunc(cronCycle, func() {
		c := database.DrawChampion()
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
