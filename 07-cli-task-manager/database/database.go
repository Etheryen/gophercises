package database

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Id     int
	Title  string
	DoneAt time.Time
}

var tasksBucket = []byte("tasks")

var db *bolt.DB

func Init(dbPath string) error {
	openedDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}

	db = openedDb

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

func GetAllTasks() ([]Task, error) {
	var allTasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			allTasks = append(allTasks, task)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return allTasks, nil
}

func CreateTask(title string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)

		id64, err := b.NextSequence()
		if err != nil {
			return err
		}

		id := int(id64)
		t := Task{Id: id, Title: title}

		jsonBytes, err := json.Marshal(t)
		if err != nil {
			return err
		}

		key := itob(id)
		return b.Put(key, jsonBytes)
	})
}

func DoTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		key := itob(id)

		jsonBytes := b.Get(key)
		if jsonBytes == nil {
			return errors.New("task not found")
		}

		var task Task

		err := json.Unmarshal(jsonBytes, &task)
		if err != nil {
			return err
		}

		task.DoneAt = time.Now()

		newJsonBytes, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(key, newJsonBytes)
	})
}

func RemoveTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		key := itob(id)
		return b.Delete(key)
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
