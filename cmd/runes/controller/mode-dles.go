package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runes/cmd/runes/database"
)

/*
Guess controller sets the entrypoint of Guess Game Mode, which will save gameID on database and render page HTML
*/
func Guess(context *gin.Context) {
	gameID := context.Param("gameID")
	database.SaveGame(context, gameID)
	var playingMode string
	if context.Request.URL.Path == "/play/guess/"+gameID {
		playingMode = "guess"
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID {
		playingMode = "mayhem"
	}
	gameDraws := database.CheckDraws(gameID, playingMode)
	context.HTML(http.StatusOK, "dle.html", gin.H{
		"Title":  "DLE",
		"Mode":   playingMode,
		"GameID": gameID,
		"Draws":  gameDraws,
	})
}

func Mayhem(context *gin.Context) {

}
