package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runes/cmd/runes/database"
	"runes/cmd/runes/game-modes/dle"
)

/*
Guess controller sets the entrypoint of Guess Game Mode, which will save gameID on database and render page HTML
*/
func DLEs(context *gin.Context) {
	gameID := context.Param("gameID")
	database.SaveGame(context, gameID)
	var playingMode string
	var championID int

	table := database.SetTable(context)
	if context.Request.URL.Path == "/play/guess/"+gameID {
		playingMode = "guess"
		championID = (*dle.DailyChampion).ID
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID {
		playingMode = "mayhem"
		championID = database.CheckGameChampion(gameID, table)
	}
	champion := dle.FindChampion(championID)
	gameDraws := database.CheckDraws(gameID, playingMode)
	for i := range gameDraws {
		var gameDraw database.Draws
		gameDraw.Champion = gameDraws[i].Champion
		gameDraw = dle.Compare(gameDraw.Champion, champion, gameDraw)
		gameDraws[i] = gameDraw
	}
	fmt.Println(gameDraws)
	context.HTML(http.StatusOK, "dle.html", gin.H{
		"Title":  "DLE",
		"Mode":   playingMode,
		"GameID": gameID,
		"Draws":  gameDraws,
	})
}

func Mayhem(context *gin.Context) {

}
