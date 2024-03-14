package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runes/cmd/runes/database"
	"runes/cmd/runes/game-modes/dle"
)

/*
MayhemDrawChampion will run before func DLEs, performing two very important actions:
Draw our champion and save this game onto DB
*/
func MayhemDrawChampion(context *gin.Context) {
	gameID := context.Param("gameID")
	champion := dle.DraftChampion()
	database.SaveGame(context, gameID, champion.ID)
	context.Next()
}

func SaveGuess(context *gin.Context) {
	gameID := context.Param("gameID")
	champion := (*dle.DailyChampion)
	database.SaveGame(context, gameID, champion.ID)
	context.Next()
}

/*
Guess controller sets the entrypoint of Guess Game Mode, which will save gameID on database and render page HTML
*/
func DLEs(context *gin.Context) {
	gameID := context.Param("gameID")
	var playingMode string
	var championID int
	table := database.SetTable(context)
	if context.Request.URL.Path == "/play/guess/"+gameID {
		playingMode = "guess"
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID {
		playingMode = "mayhem"
	}
	championID = database.CheckGameChampion(gameID, table)
	champion := dle.FindChampion(championID)
	gameDraws := database.CheckDraws(gameID, playingMode)
	for i := range gameDraws {
		var gameDraw database.Draws
		gameDraw.Champion = gameDraws[i].Champion
		gameDraw = dle.Compare(gameDraw.Champion, champion, gameDraw)
		gameDraws[i] = gameDraw
	}
	context.HTML(http.StatusOK, "dle.html", gin.H{
		"Title":         "DLE",
		"Mode":          playingMode,
		"GameID":        gameID,
		"Draws":         gameDraws,
		"ChampionsList": database.ChampionsList,
	})
}
