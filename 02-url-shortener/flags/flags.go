package flags

import (
	"flag"
)

func GetAll() (string, string) {
	yamlFile := flag.String(
		"yaml",
		"",
		"yaml file with paths and urls",
	)
	jsonFile := flag.String(
		"json",
		"",
		"json file with paths and urls",
	)
	flag.Parse()

	return *yamlFile, *jsonFile
}
