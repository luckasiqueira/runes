package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hangman(context *gin.Context) {
	context.HTML(http.StatusOK, "hangman.html", gin.H{
		"Title": "Hangman",
	})
}
