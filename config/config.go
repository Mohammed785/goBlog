package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	if err:=godotenv.Load();err!=nil{
		log.Fatalln("error loading .env file")
	}
	return os.Getenv(key)
}
