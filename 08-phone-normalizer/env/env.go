package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func Load() error {
	return godotenv.Load()
}

func GetDbInfo() (DbInfo, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	if _, err := strconv.Atoi(port); err != nil {
		return DbInfo{}, errors.New("port needs to be an integer")
	}

	return DbInfo{host, port, user, password, dbname}, nil
}
