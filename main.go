package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runes/cmd/runes/game-modes/dle"
	"runes/cmd/runes/routes"
	envdata "runes/tools/envdata"
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
	port := envdata.Env.SVPort
	err := route.Run(port)
	if err != nil {
		return
	}
}
