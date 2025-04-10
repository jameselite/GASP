package configfile

var MySQLConfig string = `
package config

import (

	"%s"
	"context"
	"fmt"
	"log"
	"sync"

	"%s"
	"github.com/joho/godotenv"
)

var (
	DBconn %s
	once   sync.Once
)

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func DBconnection() (%s, error) {
	var err error

	once.Do(func() {
		dsn := "%s"
		DBconn, err = %s
		if err != nil {
			log.Fatal("Unable to connect to database:", err)
		}
		fmt.Println("Database connected successfully")
	})

	return DBconn, err
}
`