package hangman

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"math/rand"
	"net/http"
	"runes/cmd/runes/database"
	"strings"
)

var HangmanChampion = &c
var c = database.CheckHangmanChampion()

/*
POST worker
*/
func PlayHangman(context *gin.Context) {
	drawLetter := strings.ToUpper(context.PostForm("draw"))
	c := strings.ToUpper(*HangmanChampion)
	for index, letter := range c {
		if drawLetter == string(letter) {
			context.HTML(http.StatusOK, "hangman-dynamics.html", gin.H{
				"Index": index,
			})
			break
		}
	}
}

func DraftHangmanChampion() {
	cronCycle := "0 0 * * *"
	job := cron.New()
	job.AddFunc(cronCycle, func() {
		i := rand.Intn(len((*database.ChampionsList)))
		championID := (*database.ChampionsList)[i].Champion.ID
		database.SaveHangmanChampion(championID)
		for i := range *database.ChampionsList {
			if championID == (*database.ChampionsList)[i].Champion.ID {
				c = (*database.ChampionsList)[i].Champion.Name
				break
			}
		}
	})
	job.Start()
	job.Run()
}
