package quiz

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type QuestionAnswer struct {
	question string
	answer   string
}

func ParseRecords(
	records [][]string,
	shouldShuffle bool,
) ([]QuestionAnswer, error) {
	var result []QuestionAnswer

	for _, record := range records {
		if len(record) != 2 {
			return nil, errors.New("records should have 2 columns")
		}
		result = append(result, QuestionAnswer{
			question: record[0],
			answer:   record[1],
		})
	}

	if shouldShuffle {
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)
		rng.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
	}

	return result, nil
}

func askQuestionCheckAnswer(
	questionAnswer QuestionAnswer,
	number int,
	resultChan chan bool,
) {
	fmt.Printf("Problem #%v: %v = ", number, questionAnswer.question)

	var userAnswer string
	fmt.Scanln(&userAnswer)
	cleanedAnswer := strings.ToLower(strings.Trim(userAnswer, " "))
	cleanedCorrectAnswer := strings.ToLower(
		strings.Trim(questionAnswer.answer, " "),
	)

	resultChan <- cleanedAnswer == cleanedCorrectAnswer
}

func announce(duration time.Duration, questionsAmount int) {
	fmt.Printf(
		"The quiz is about to start, you will have %v to answer %v questions, good luck! Hit enter when ready...",
		duration,
		questionsAmount,
	)
	fmt.Scanln()
	fmt.Println("GOOO!!!")
}

func Start(
	questionsAnswers []QuestionAnswer,
	timerDuration time.Duration,
) int {
	announce(timerDuration, len(questionsAnswers))

	timer := time.NewTimer(timerDuration)
	resultChan := make(chan bool)
	score := 0

	for i, questionAnswer := range questionsAnswers {
		go askQuestionCheckAnswer(questionAnswer, i+1, resultChan)

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			return score
		case isCorrect := <-resultChan:
			if isCorrect {
				score++
			}
		}
	}

	return score
}
