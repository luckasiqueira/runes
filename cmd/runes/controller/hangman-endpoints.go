package controller

import (
	"github.com/gin-gonic/gin"
)

func HangmanDraws(context *gin.Context) {
	drawChampion := context.PostForm("draw")
}
