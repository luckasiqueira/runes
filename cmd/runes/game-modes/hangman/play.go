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

type HangmanDraws struct {
	Letter string
	Found  bool
}

/*
POST worker
*/
func PlayHangman(context *gin.Context) {
	draw := strings.ToUpper(context.PostForm("draw"))
	champion := (*HangmanChampion)
	gameID := context.Param("gameID")
	database.SaveHangmanGameDraws(gameID, draw)
	draws := database.CheckHangmanDraws(gameID)
	for index, letter := range champion {
		if draw == string(letter) {
			context.HTML(http.StatusOK, "hangman-dynamics.html", gin.H{
				"Index": index,
				"Draws": draws,
			})
			break
		}
	}
}

func Compare(draws []string, champion []string) []HangmanDraws {
	var gameDraws []HangmanDraws
	for i := range draws {
		var draw HangmanDraws
		for j := range champion {
			if draws[i] == champion[j] {
				draw.Letter = champion[j]
				draw.Found = true
				gameDraws = append(gameDraws, draw)
			}
		}
	}
	return gameDraws
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
				c = append(c, (*database.ChampionsList)[i].Champion.Name)
				break
			}
		}
	})
	job.Start()
	job.Run()
}
