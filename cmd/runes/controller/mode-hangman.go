package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runes/cmd/runes/database"
	"runes/tools/envdata"
	"strings"
)

func Hangman(context *gin.Context) {
	letters := []string{}
	for i := 'a'; i <= 'z'; i++ {
		letters = append(letters, strings.ToUpper(string(i)))
	}
	c := "Viego"
	var champion []string
	for _, i := range c {
		champion = append(champion, strings.ToUpper(string(i)))
	}
	context.HTML(http.StatusOK, "hangman.html", gin.H{
		"Title":    "Hangman",
		"Champion": champion,
		"Alphabet": letters,
	})
}

func SaveGameHangman(context *gin.Context) {
	table := envdata.Env.TBHangman
	gameID := context.Param("gameID")
	database.SaveHangman(gameID, table)
	context.Next()
}
