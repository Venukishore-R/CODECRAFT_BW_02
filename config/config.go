package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DBUSERNAME string
	DBPASSWORD string
	DBHOST     string
	DBPORT     string
	DBNAME     string
}

func LoadEnv() *Config {
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if userName == "" || password == "" || host == "" || port == "" || dbName == "" {
		log.Fatalf("Missing environment variables")
	}

	config := &Config{
		DBUSERNAME: userName,
		DBPASSWORD: password,
		DBHOST:     host,
		DBPORT:     port,
		DBNAME:     dbName,
	}

	return config
}
