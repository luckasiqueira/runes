package controller

import (
	"github.com/gin-gonic/gin"
	"runes/cmd/runes/game-modes/dle"
	"strings"
)

/*
DLEDraws collects user's shots to draw var, which will be converted as uppercase
*/
func DLEDraws(context *gin.Context) {
	draw := strings.ToUpper(context.PostForm("draw"))
	dle.PlayDLE(context, draw)
}
