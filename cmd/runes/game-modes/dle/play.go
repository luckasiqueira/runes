package dle

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"net/http"
	"runes/cmd/runes/database"
	"strings"
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
	var gameDraw database.Draws
	var gameDraws []database.Draws
	table := database.SetTable(context)
	var playingMode string
	if context.Request.URL.Path == "/try/guess/"+gameID {
		championID = (*dailyChampion).ID
		playingMode = "guess"
	} else if context.Request.URL.Path == "/try/mayhem/"+gameID {
		championID = database.CheckGameChampion(gameID, table)
		playingMode = "mayhem"
	}
	drawChampion := database.CheckChampionDrawed(draw)
	gameDraw.Champion = drawChampion
	champion := findChampion(championID)
	gameDraw = compare(championID, drawChampion, champion, gameDraw)
	gameDraws = append(gameDraws, gameDraw)
	context.HTML(http.StatusOK, "dle-dynamics.html", gin.H{
		"Draws": gameDraws,
	})
	go database.SaveDraw(gameID, playingMode, drawChampion.ID)
}

/*
findChampion gets the given championID and loops over ChampionsList, trying to match this ID with a defined champion
Once a champion is found, it's saved and returned by 'champion' var
*/
func findChampion(championID int) database.ChampionLOL {
	var champion database.ChampionLOL
	for i := range *database.ChampionsList {
		if championID == (*database.ChampionsList)[i].Champion.ID {
			champion = (*database.ChampionsList)[i].Champion
			break
		}
	}
	return champion
}

/*
compare compares all characteristics for the given champion drawed with the defined champion for this game.
If a characteristic is correct, it will set a Status as true, which will be used to indicate to player if that shot is correct, partially correct or wrong.
*/
func compare(championID int, drawChampion, champion database.ChampionLOL, gameDraw database.Draws) database.Draws {
	if drawChampion.ID == championID {
		gameDraw.Won = true
	} else {
		if drawChampion.Gender == champion.Gender {
			gameDraw.Status.GenderFound = true
		}
		if drawChampion.Role == champion.Role {
			gameDraw.Status.RoleFound = true
		} else if isContained(drawChampion.Role, champion.Role) {
			gameDraw.Status.RolePartial = true
		}
		if drawChampion.Race == champion.Race {
			gameDraw.Status.RaceFound = true
		} else if isContained(drawChampion.Race, champion.Race) {
			gameDraw.Status.RacePartial = true
		}
		if drawChampion.Resource == champion.Resource {
			gameDraw.Status.ResourceFound = true
		}
		if drawChampion.Range == champion.Range {
			gameDraw.Status.RangeFound = true
		} else if isContained(drawChampion.Range, champion.Range) {
			gameDraw.Status.RangePartial = true
		}
		if drawChampion.Region == champion.Region {
			gameDraw.Status.RegionFound = true
		} else if isContained(drawChampion.Region, champion.Region) {
			gameDraw.Status.RegionPartial = true
		}
		if drawChampion.Release == champion.Release {
			gameDraw.Status.ReleaseFound = true
		} else if drawChampion.Release > champion.Release {
			gameDraw.Status.ReleaseUp = true
		} else if drawChampion.Release < champion.Release {
			gameDraw.Status.ReleaseDown = true
		}
	}
	return gameDraw
}

/*
isContained evaluates characteristics for both champions (user given and game set) in order to identify which are equal
*/
func isContained(d, c string) bool {
	draw := strings.Split(strings.ToUpper(d), ", ")
	champion := strings.Split(strings.ToUpper(c), ", ")
	for _, wordA := range draw {
		for _, wordB := range champion {
			if wordA == wordB {
				return true
			}
		}
	}
	return false
}

/*
DraftDailyChampion runs every 00:00h (here, set as 12), when it Draws a new champion from DB and saves onto dailyChampion pointer
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
