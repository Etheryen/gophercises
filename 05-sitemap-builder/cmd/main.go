package main

import (
	"05-sitemap-builder/internal/flags"
	"05-sitemap-builder/internal/sitemap"
	"fmt"
	"log"
)

func main() {
	url, depth := flags.GetAll()

	if url == "" {
		log.Fatalln("Error reading flags: you need to provide a url")
	}

	sitemap, err := sitemap.Build(url, depth)
	if err != nil {
		log.Fatalln("Error building sitemap:", err)
	}

	fmt.Println(string(sitemap))
}
