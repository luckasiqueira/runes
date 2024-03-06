package dle

import (
	"github.com/robfig/cron/v3"
	"runes/cmd/runes/database"
)

var dailyChampion *database.ChampionLOL

/*
 */
func PlayGuess() {
	dailyDraw()
}

/*
dailyDraw runs every 00:00 (here, set as 12), when it Draws a new champion from DB and saves onto dailyChampion pointer
This pointer will be used to compare user shots fast, since no DB comparision will be needed
*/
func dailyDraw() {
	cronCycle := "* * * * *"
	job := cron.New()
	job.AddFunc(cronCycle, func() {
		c := database.DrawChampion()
		dailyChampion = &c
	})
	job.Start()
	job.Run()
}

func PlayMayhem() {

}
