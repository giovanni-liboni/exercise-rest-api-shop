package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func ConnectDB(config *Config) *sqlx.DB {
	//	// Build the connection string from the config
	dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName)

	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
