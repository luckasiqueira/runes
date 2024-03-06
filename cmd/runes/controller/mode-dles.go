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
	context.HTML(http.StatusOK, "dle.html", gin.H{
		"Title":  "DLE",
		"Mode":   "guess",
		"GameID": gameID,
	})
}

func Mayhem(context *gin.Context) {

}
