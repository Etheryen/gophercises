package timer

import (
	"01-quiz/state"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

const DEFAULT_TIME = 30 * time.Second

func announce(duration time.Duration) error {
	fmt.Println(
		"The quiz is about to start, you will have",
		duration,
		"to answer",
		state.QuestionsTotal,
		"questions, good luck! Hit enter when ready...",
	)
	for {
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			return err
		}
		if key == keyboard.KeyEnter {
			break
		}
	}
	fmt.Println("GOOO!!!")
	return nil
}

func stopAfter(duration time.Duration) {
	time.Sleep(duration)
	fmt.Println("\nTime is up!")
	state.PrintResults()
	os.Exit(0)
}

func Start(duration time.Duration) error {
	err := announce(duration)
	if err != nil {
		return err
	}
	go stopAfter(duration)
	return nil
}
