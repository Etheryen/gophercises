package main

import (
	"05-sitemap-builder/internal/flags"
	"05-sitemap-builder/internal/sitemap"
	"fmt"
	"log"
)

func main() {
	// TODO: use depth and BFS algorythm
	url, _ := flags.GetAll()

	if url == "" {
		log.Fatalln("Error reading flags: you need to provide a url")
	}

	sitemap, err := sitemap.Build(url)
	if err != nil {
		log.Fatalln("Error building sitemap:", err)
	}

	prettyPrint(sitemap, "")
}

func prettyPrint(sitemap sitemap.SiteNode, padding string) {
	fmt.Printf("%v%v\n", padding, sitemap.Location)
	for _, child := range sitemap.Children {
		prettyPrint(child, padding+"  ")
	}
}
