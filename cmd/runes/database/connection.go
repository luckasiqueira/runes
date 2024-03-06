package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Connect() -> error while loading .env file.")
	}
	data := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", data)
	db.SetConnMaxLifetime(time.Minute * 5)
	return db
}

/*
SetTable checks the given URL Path to identify game mode.
When game mode is identified, atributes a value set on .env file to table var, which is returned
*/
func SetTable(context *gin.Context) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Connect() -> error while loading .env file.")
	}
	var table string
	gameID := context.Param("gameID")
	if context.Request.URL.Path == "/play/guess/"+gameID || context.Request.URL.Path == "/try/guess/"+gameID {
		table = os.Getenv("TB_GUESS")
	} else if context.Request.URL.Path == "/play/mayhem/"+gameID || context.Request.URL.Path == "/try/mayhem/"+gameID {
		table = os.Getenv("TB_MAYHEM")
	}
	return table
}
