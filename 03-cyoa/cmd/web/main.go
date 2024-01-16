package main

import (
	"03-cyoa/flags"
	"03-cyoa/handlers"
	"03-cyoa/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	file, port := flags.GetAll()
	portString := ":" + strconv.Itoa(port)

	storyJson, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading story file", err)
	}

	storyMap, err := utils.ParseStoryJSON(storyJson)
	if err != nil {
		log.Fatalln("Error parsing story json:", err)
	}

	handler := handlers.NewHandler(storyMap)

	fmt.Println("Listening on port", portString)
	err = http.ListenAndServe(portString, handler)
	if err != nil {
		log.Fatalln("HTTP server error:", err)
	}
}
