package main

import (
	"01-quiz/flags"
	"01-quiz/quiz"
	"01-quiz/utils"
	"encoding/csv"
	"fmt"
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

	score := quiz.Start(parsed, timerDuration)

	perentage := 100 * score / len(parsed)
	fmt.Printf(
		"Score: %v/%v, that's %v%%!\n",
		score,
		len(parsed),
		perentage,
	)
}
