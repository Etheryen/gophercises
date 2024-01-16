package flags

import (
	"flag"
)

const (
	DEFAULT_FILE string = "gopher.json"
	DEFAULT_PORT int    = 8080
)

func GetAll() (string, int) {
	file := flag.String(
		"file",
		DEFAULT_FILE,
		"json file with story arcs",
	)
	port := flag.Int("port", DEFAULT_PORT, "http server port")

	flag.Parse()

	return *file, *port
}
