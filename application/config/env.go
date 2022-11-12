package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	mongoIRU := strings.TrimSpace(os.Getenv("MONGOURI"))
	if len(mongoIRU) == 0 {

		err := godotenv.Load()

		if err != nil {
			log.Fatalln("error .env")
		}

		mongoIRU = os.Getenv("MONGOURI")
	}
	return mongoIRU
}
