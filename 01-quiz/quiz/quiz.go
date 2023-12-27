package quiz

import (
	"01-quiz/state"
	"01-quiz/timer"
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

func askQuestionCheckAnswer(questionAnswer QuestionAnswer, number int) bool {
	fmt.Printf("Problem #%v: %v = ", number, questionAnswer.question)

	var userAnswer string

	fmt.Scanln(&userAnswer)

	cleanedAnswer := strings.ToLower(strings.Trim(userAnswer, " "))

	return cleanedAnswer == questionAnswer.answer
}

func Start(
	questionsAnswers []QuestionAnswer,
	timerDuration time.Duration,
) error {
	state.QuestionsTotal = len(questionsAnswers)

	err := timer.Start(timerDuration)
	if err != nil {
		return err
	}

	for i, questionAnswer := range questionsAnswers {
		isCorrect := askQuestionCheckAnswer(questionAnswer, i+1)
		if isCorrect {
			state.Score++
		}
	}

	return nil
}
