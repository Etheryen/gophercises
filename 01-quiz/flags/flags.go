package flags

import (
	"flag"
	"time"
)

const DEFAULT_FILE string = "problems.csv"

func GetAll() (string, time.Duration, bool) {
	file := flag.String(
		"file",
		DEFAULT_FILE,
		"csv file with quiz content",
	)
	timerSeconds := flag.Int("time", 30, "quiz time (in seconds)")
	shouldShuffle := flag.Bool("shuffle", false, "shuffle the questions")

	flag.Parse()

	return *file, time.Second * time.Duration(*timerSeconds), *shouldShuffle
}
