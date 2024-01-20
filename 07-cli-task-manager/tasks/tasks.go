package tasks

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	Id     int
	Title  string
	DoneAt time.Time
}

var tasksDb []Task = []Task{
	{1, "yooo", time.Now()},
	{Id: 2, Title: "bruh"},
}

func Add(title string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}

	newId := tasksDb[len(tasksDb)-1].Id + 1

	tasksDb = append(tasksDb, Task{Id: newId, Title: title})
	fmt.Printf("Added \"%v\" to your task list.\n", title)

	return nil
}

func List() {
	undone := filterUndone(tasksDb)

	if len(undone) == 0 {
		fmt.Println("You have no uncompleted tasks.")
		return
	}

	fmt.Println("You have the following tasks:")

	for i, t := range undone {
		fmt.Printf("%v. %v\n", i+1, t.Title)
	}
}

func filterUndone(taskList []Task) []Task {
	var undone []Task

	for _, t := range taskList {
		if t.DoneAt.IsZero() {
			undone = append(undone, t)
		}
	}

	return undone
}

// TODO: check if pointer works
func Do(num int) error {
	undone := filterUndone(tasksDb)

	if num < 1 || num > len(undone) {
		return errors.New("invalid task number")
	}

	id := undone[num-1].Id
	taskDone, err := getTask(tasksDb, id)
	if err != nil {
		return err
	}
	doTask(tasksDb, id)

	fmt.Printf("You have completed the \"%v\" task.\n", taskDone.Title)

	return nil
}

func getTask(taskList []Task, id int) (Task, error) {
	for _, t := range taskList {
		if t.Id == id {
			return t, nil
		}
	}

	return Task{}, errors.New("task not found")
}

func doTask(taskList []Task, id int) {
	for i, t := range taskList {
		if t.Id == id {
			taskList[i].DoneAt = time.Now()
		}
	}
}

func Remove(num int) error {
	undone := filterUndone(tasksDb)
	if num < 1 || num > len(undone) {
		return errors.New("invalid task number")
	}

	removed := undone[num-1]
	tasksDb = withoutTask(tasksDb, num-1)
	fmt.Printf("You have deleted the \"%v\" task.\n", removed.Title)

	return nil
}

func withoutTask(taskList []Task, id int) []Task {
	result := make([]Task, len(taskList)-1)

	for _, t := range taskList {
		if t.Id != id {
			result = append(result, t)
		}
	}

	return result
}

func Completed() {
	completed := filterCompletedToday(tasksDb)
	if len(completed) == 0 {
		fmt.Println("You haven't completed any tasks today yet.")
		return
	}

	fmt.Println("You have finished the following tasks today:")
	for _, t := range completed {
		fmt.Printf("- %v\n", t.Title)
	}
}

func filterCompletedToday(taskList []Task) []Task {
	var completed []Task

	now := time.Now()

	for _, t := range taskList {
		if isSameDate(t.DoneAt, now) {
			completed = append(completed, t)
		}
	}

	return completed
}

func isSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
