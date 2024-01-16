package main

import (
	"04-html-link-parser/internal/parsing"
	"fmt"
	"log"
	"os"
)

func main() {
	htmlData, err := os.ReadFile("internal/parsing/test_data/ex3.html")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	parsed, err := parsing.ParseLinks(htmlData)
	if err != nil {
		log.Fatalln("Error parsing links:", err)
	}

	fmt.Println(parsed)
}
