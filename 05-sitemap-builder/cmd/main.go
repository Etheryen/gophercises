package main

import (
	"05-sitemap-builder/internal/flags"
	"05-sitemap-builder/internal/sitemap"
	"fmt"
	"log"
	"time"
)

func main() {
	url, depth := flags.GetAll()

	if url == "" {
		log.Fatalln("Error reading flags: you need to provide a url")
	}

	tStart := time.Now()

	sitemap, err := sitemap.Build(url, depth)
	if err != nil {
		log.Fatalln("Error building sitemap:", err)
	}
	tDelta := time.Since(tStart)

	fmt.Println(string(sitemap))
	fmt.Println("Took", tDelta)
}
