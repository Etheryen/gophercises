package state

import "fmt"

var (
	Score          int
	QuestionsTotal int
)

func PrintResults() {
	perentage := 100 * Score / QuestionsTotal
	fmt.Printf(
		"Score: %v/%v, that's %v%%!\n",
		Score,
		QuestionsTotal,
		perentage,
	)
}
