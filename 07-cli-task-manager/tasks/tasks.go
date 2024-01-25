package tasks

import (
	"errors"
	"fmt"
	"strings"
	"task/database"
	"time"
)

func Add(title string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}

	err := database.CreateTask(title)
	if err != nil {
		return err
	}

	fmt.Printf("Added \"%v\" to your task list.\n", title)
	return nil
}

func List() error {
	allTasks, err := database.GetAllTasks()
	if err != nil {
		return err
	}

	undone := filterUndone(allTasks)

	if len(undone) == 0 {
		fmt.Println("You have no uncompleted tasks.")
		return nil
	}

	fmt.Println("You have the following tasks:")

	for i, t := range undone {
		fmt.Printf("%v. %v\n", i+1, t.Title)
	}

	return nil
}

func filterUndone(taskList []database.Task) []database.Task {
	var undone []database.Task

	for _, t := range taskList {
		if t.DoneAt.IsZero() {
			undone = append(undone, t)
		}
	}

	return undone
}

func Do(num int) error {
	allTasks, err := database.GetAllTasks()
	if err != nil {
		return err
	}

	undone := filterUndone(allTasks)

	if num < 1 || num > len(undone) {
		return errors.New("invalid task number")
	}

	task := undone[num-1]
	err = database.DoTask(task.Id)
	if err != nil {
		return err
	}

	fmt.Printf("You have completed the \"%v\" task.\n", task.Title)

	return nil
}

func Remove(num int) error {
	allTasks, err := database.GetAllTasks()
	if err != nil {
		return err
	}

	undone := filterUndone(allTasks)

	if num < 1 || num > len(undone) {
		return errors.New("invalid task number")
	}

	task := undone[num-1]
	err = database.RemoveTask(task.Id)
	if err != nil {
		return err
	}

	fmt.Printf("You have deleted the \"%v\" task.\n", task.Title)

	return nil
}

func Completed() error {
	allTasks, err := database.GetAllTasks()
	if err != nil {
		return err
	}

	completed := filterCompletedToday(allTasks)

	if len(completed) == 0 {
		fmt.Println("You haven't completed any tasks today yet.")
		return nil
	}

	fmt.Println("You have finished the following tasks today:")
	for _, t := range completed {
		fmt.Printf("- %v\n", t.Title)
	}

	return nil
}

func filterCompletedToday(taskList []database.Task) []database.Task {
	var completed []database.Task

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
