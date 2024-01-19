package flags

import (
	"flag"
)

func GetAll() (string, int) {
	url := flag.String(
		"url",
		"",
		"url of the website to map the pages of",
	)
	depth := flag.Int("depth", 3, "maximum depth of links to follow")

	flag.Parse()

	return *url, *depth
}
