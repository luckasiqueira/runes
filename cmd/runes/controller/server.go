package controller

import (
	"github.com/gin-gonic/gin"
	"runes/tools/envdata"
)

/*
ServerInfo sets all server details and headers
*/
func ServerInfo(context *gin.Context) {
	context.Header("Server-Node", envdata.Env.HeaderNode)
	context.Next()
}
