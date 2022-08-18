package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Connect() {
	DB_CREDENTIALS := goDotEnvVariable("DB_CREDENTIALS")
	d, err := gorm.Open("mysql", DB_CREDENTIALS)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
