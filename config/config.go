package config
import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Jwt      JwtConfig
	Timezone string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type ServerConfig struct {
	Port int
}

type JwtConfig struct {
	Secret string
}

func loadFromFile(filename string) {
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		err := godotenv.Load(filename)
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	} else {
		log.Fatal("Error loading .env file", err)
	}
}

func LoadConfig(params... string) *Config {
	switch len(params) {
	case 1:
		loadFromFile(params[0])
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error converting port to int", err)
	}

	portDb := os.Getenv("DB_PORT")
	if portDb == "" {
		portDb = "3306"
	}
	portDbInt, err := strconv.Atoi(portDb)
	if err != nil {
		log.Fatal("Error converting database port to int", err)
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     portDbInt,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},
		Server: ServerConfig{
			Port: portInt,
		},
		Jwt: JwtConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Timezone: os.Getenv("TIMEZONE"),
	}
	return config
}

