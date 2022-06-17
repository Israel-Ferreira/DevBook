package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl   string
	Porta    int
	HashKey  []byte
	BlockKey []byte
)

func LoadEnv() {
	var erro error

	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado")
	}

	ApiUrl = os.Getenv("API_URL")

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))

	if erro != nil {
		log.Fatalln(erro)
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
