package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Database     DatabaseConfig
	Server       ServerConfig
	Jwt          JwtConfig
	Timezone     string
	StripeAPIKey string
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

// mysql://mysql:a321141c10fe3bba@dokku-mysql-exercise-rest-api-shop-db:3306/exercise_rest_api_shop_db
// Set env variables DB_HOST from the DATABASE_URL environment variable

func parseDatabaseUrl() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return
	}
	u, _ := url.Parse(databaseUrl)

	err := os.Setenv("DB_HOST", u.Hostname())
	if err != nil {
		log.Errorln("Error setting DB_HOST", err)
	}
	err = os.Setenv("DB_PORT", u.Port())
	if err != nil {
		log.Errorln("Error setting DB_PORT", err)
	}

	err = os.Setenv("DB_USER", u.User.Username())
	if err != nil {
		log.Errorln("Error setting DB_USER", err)
	}

	if pass, ok := u.User.Password(); ok {
		err = os.Setenv("DB_PASSWORD", pass)
		if err != nil {
			log.Errorln("Error setting DB_PASSWORD", err)
		}
	}

	err = os.Setenv("DB_NAME", u.Path[1:])
	if err != nil {
		log.Errorln("Error setting DB_NAME", err)
	}
}

func LoadConfig(params ...string) *Config {
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

	// Set env variables for the database if any host is set. It will load the database url
	// from the DATABASE_URL environment variable
	if os.Getenv("DB_HOST") == "" {
		parseDatabaseUrl()
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
		Timezone:     os.Getenv("TIMEZONE"),
		StripeAPIKey: os.Getenv("STRIPE_API_KEY"),
	}
	return config
}
