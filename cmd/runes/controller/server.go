package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
ServerInfo sets all server details and headers
*/
func ServerInfo(context *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ServerInfo() -> error while loading .env file")
	}
	headerNode := os.Getenv("H_NODE")
	context.Header("Server-Node", headerNode)
	context.Next()
}
