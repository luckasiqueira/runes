package controller

import (
	"github.com/gin-gonic/gin"
	"runes/cmd/runes/game-modes/hangman"
)

func HangmanDraws(context *gin.Context) {
	hangman.PlayHangman(context)
}
