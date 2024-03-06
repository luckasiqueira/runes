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
		dailyDraw()
		champion = *(dailyChampion)
	} else if context.Request.URL.Path == "/try/mayhem/"+gameID {
		championID = database.CheckGameChampion(gameID, table)
	}
	fmt.Println(championID, champion)
}

/*
dailyDraw runs every 00:00 (here, set as 12), when it Draws a new champion from DB and saves onto dailyChampion pointer
This pointer will be used to compare user shots fast, since no DB comparision will be needed
*/
func draftDailyChampion() {
	cronCycle := "* * * * *"
	job := cron.New()
	job.AddFunc(cronCycle, func() {
		c := database.DrawChampion()
		dailyChampion = &c
	})
	job.Start()
	job.Run()
}
