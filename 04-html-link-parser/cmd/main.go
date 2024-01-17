package main

import (
	"04-html-link-parser/internal/parsing"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	htmlData, err := os.ReadFile("internal/parsing/test_data/ex3.html")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	r := bytes.NewReader(htmlData)
	parsed, err := parsing.ParseLinks(r)
	if err != nil {
		log.Fatalln("Error parsing links:", err)
	}

	fmt.Println(parsed)
}
