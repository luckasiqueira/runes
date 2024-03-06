package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"runes/cmd/runes/game-modes/dle"
	"runes/cmd/runes/routes"
)

/*
main starts our entire application, setting an WebServer with Gin-Gonic.
gin.Default starts all Gin basic configuration
LoadHTMLGlob allows to user our .html files to render page's content
StatictFS redirect /assets/ requests to website/assets directory, in order to allow static files (.css, .js) delivery
Router starts our routes settings
Run ... dah...
*/
func main() {
	route := gin.Default()
	dle.DraftDailyChampion()
	route.LoadHTMLGlob("website/*html")
	route.StaticFS("/assets/", http.Dir("website/assets"))
	routes.Router(&route.RouterGroup)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("main -> error while loading .env file")
	}
	port := os.Getenv("SV_PORT")
	err = route.Run(port)
	if err != nil {
		return
	}
}
