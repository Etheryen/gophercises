package main

import (
	"01-quiz/flags"
	"01-quiz/quiz"
	"01-quiz/state"
	"01-quiz/utils"
	"encoding/csv"
	"log"
	"strings"
)

func main() {
	fileName, timerDuration, shouldShuffle := flags.GetAll()

	file, err := utils.FileToString(fileName)
	if err != nil {
		log.Fatalln("Error reading file:", err.Error())
	}

	r := csv.NewReader(strings.NewReader(file))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Error parsing csv:", err.Error())
	}

	parsed, err := quiz.ParseRecords(records, shouldShuffle)
	if err != nil {
		log.Fatalln("Error parsing records:", err.Error())
	}

	err = quiz.Start(parsed, timerDuration)
	if err != nil {
		log.Fatalln("Error during quiz:", err.Error())
	}

	state.PrintResults()
}
