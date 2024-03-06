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
	route.GET("/", controller.Index)
	play := route.Group("/play")
	{
		play.GET("/guess/", controller.RedirectGameID)
		play.GET("/guess/:gameID", controller.CheckGameIsValid, controller.Guess)
		play.GET("/mayhem/", controller.RedirectGameID)
		play.GET("/mayhem/:gameID", controller.CheckGameIsValid, controller.Mayhem)
	}
	try := route.Group("/try")
	{
		try.POST("/guess/:gameID", controller.DLEDraws)
		try.POST("/mayhem/:gameID", controller.DLEDraws)
	}
}
