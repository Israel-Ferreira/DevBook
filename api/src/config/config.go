package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConexaoDbString string = ""
	Porta           int    = 8090

	SecretKey []byte
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Println("erro ao carregar o arquivo .env")
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUrl := os.Getenv("DB_URL")

	secretKey := os.Getenv("SECRET_KEY")

	ConexaoDbString = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbUrl,
		dbName,
	)

	SecretKey = []byte(secretKey)

}
