package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"runes/tools/game-ID"
)

/*
RedirectGameID generates an gameID and redirect /play/<game>/ to that new gameID, to start a new game
*/
func RedirectGameID(context *gin.Context) {
	gameID := game_ID.IDGen().String()
	context.Redirect(http.StatusMovedPermanently, context.FullPath()+gameID)
}

/*
CheckGameIsValid checks if the given param gameID is valid and, case it's not, generates a new one and redirects to it
*/
func CheckGameIsValid(context *gin.Context) {
	gameID := context.Param("gameID")
	_, err := uuid.Parse(gameID)
	if err != nil {
		newID := game_ID.IDGen().String()
		context.Redirect(http.StatusMovedPermanently, newID)
		context.Abort()
		return
	}
	context.Next()
}
