package main

import (
	"log"
	"path/filepath"
	"task/cmd"
	"task/database"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln("Error locating home dir:", err)
	}

	dbPath := filepath.Join(home, "tasks.db")

	err = database.Init(dbPath)
	if err != nil {
		log.Fatalln("Error initializing database:", err)
	}

	cmd.Execute()
}
