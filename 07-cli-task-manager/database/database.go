package database

import (
	"log"
	"task/tasks"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")

func Update(taskList []tasks.Task) {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		// _, err := tx.CreateBucketIfNotExists(tasksBucket)
		return nil
	})
}
