package envdata

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvInfo struct {
	DBHost          string
	DBName          string
	DBUser          string
	DBPass          string
	DBPort          string
	SVPort          string
	TBChampions     string
	TBMayhem        string
	TBMayhemDraws   string
	TBGuess         string
	TBGuessDraws    string
	TBGuessChampion string
	HeaderNode      string
}

var Env = envLoader()

func envLoader() *EnvInfo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("envLoader() -> error while loading .env file")
	}
	e := &EnvInfo{
		DBHost:          os.Getenv("DB_HOST"),
		DBName:          os.Getenv("DB_NAME"),
		DBUser:          os.Getenv("DB_USER"),
		DBPass:          os.Getenv("DB_PASS"),
		DBPort:          os.Getenv("DB_PORT"),
		SVPort:          os.Getenv("SV_PORT"),
		TBChampions:     os.Getenv("TB_CHAMPIONS"),
		TBMayhem:        os.Getenv("TB_MAYHEM"),
		TBMayhemDraws:   os.Getenv("TB_MAYHEM_DRAWS"),
		TBGuess:         os.Getenv("TB_GUESS"),
		TBGuessDraws:    os.Getenv("TB_GUESS_DRAWS"),
		TBGuessChampion: os.Getenv("TB_GUESS_CHAMPION"),
		HeaderNode:      os.Getenv("H_NODE"),
	}
	return e
}
