package hangman

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PlayHangman(context *gin.Context) {
	drawLetter := strings.ToUpper(context.PostForm("draw"))
	c := strings.ToUpper("Viego")
	for index, letter := range c {
		if drawLetter == string(letter) {
			context.HTML(http.StatusOK, "hangman-dynamics.html", gin.H{
				"Index": index,
			})
			break
		}
	}
}
