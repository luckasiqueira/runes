package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"runes/tools/envdata"
	"time"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Connect() -> error while loading .envdata file.")
	}
	data := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		envdata.Env.DBUser,
		envdata.Env.DBPass,
		envdata.Env.DBHost,
		envdata.Env.DBPort,
		envdata.Env.DBName,
	)
	db, err := sql.Open("mysql", data)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db
}

/*
SetTable checks the given URL Path to identify game mode.
When game mode is identified, atributes a value set on .envdata file to table var, which is returned
*/
func SetTable(context *gin.Context) string {
	var table string
	gameID := context.Param("gameID")
	if context.Request.URL.Path == "/play/guess/"+gameID || context.Request.URL.Path == "/try/guess/"+gameID {
		table = envdata.Env.TBGuess
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID || context.Request.URL.Path == "/try/mayhem/"+gameID {
		table = envdata.Env.TBMayhem
	}
	return table
}
