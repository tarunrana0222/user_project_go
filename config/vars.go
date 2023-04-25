package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port       string
	Mongo_Uri  string
	Jwt_Secret string
	DB_Name    string
	EnvLoaded  = false
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port = os.Getenv("PORT")
	Mongo_Uri = os.Getenv("MONGO_URI")
	Jwt_Secret = os.Getenv("JWT_SECRET")
	DB_Name = os.Getenv("DB_NAME")

	EnvLoaded = true

	fmt.Println("Loaded Env")
}
