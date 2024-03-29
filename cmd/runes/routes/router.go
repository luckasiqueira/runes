package routes

import (
	"github.com/gin-gonic/gin"
	"runes/cmd/runes/controller"
)

/*
Router defines all our application routes, also sets their controllers and validations
Group allows set child routes, from a primary one
*/
func Router(route *gin.RouterGroup) {
	route.GET("/", controller.ServerInfo, controller.Index)
	play := route.Group("/play")
	{
		play.GET("/guess/", controller.RedirectGameID)
		play.GET("/guess/:gameID", controller.ServerInfo, controller.CheckGameIsValid, controller.SaveGuess, controller.DLEs)
		play.GET("/mayhem/", controller.RedirectGameID)
		play.GET("/mayhem/:gameID", controller.ServerInfo, controller.CheckGameIsValid, controller.MayhemDrawChampion, controller.DLEs)
		play.GET("/hangman/:gameID", controller.ServerInfo, controller.CheckGameIsValid, controller.SaveGameHangman, controller.Hangman)
	}
	try := route.Group("/try")
	{
		try.POST("/guess/:gameID", controller.DLEDraws)
		try.POST("/mayhem/:gameID", controller.DLEDraws)
		try.POST("/hangman/:gameID", controller.HangmanDraws)
	}
}
