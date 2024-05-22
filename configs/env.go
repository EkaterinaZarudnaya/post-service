package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv() string {
	if dbHost := os.Getenv("DB_HOST"); dbHost == "" {
		err := godotenv.Load()
		if err != nil {
			err = godotenv.Load("../.env")
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		}
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPwd + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Europe/Kyiv search_path=posts"
	return dsn
}
