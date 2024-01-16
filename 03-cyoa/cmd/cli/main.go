package main

import (
	"03-cyoa/cli"
	"03-cyoa/flags"
	"03-cyoa/utils"
	"log"
	"os"
)

func main() {
	file, _ := flags.GetAll()

	storyJson, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading story file", err)
	}

	storyMap, err := utils.ParseStoryJSON(storyJson)
	if err != nil {
		log.Fatalln("Error parsing story json:", err)
	}

	cli.Start(storyMap, "intro")
}
