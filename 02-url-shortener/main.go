package main

import (
	"02-url-shortener/flags"
	"02-url-shortener/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yamlFile, jsonFile := flags.GetAll()

	if jsonFile != "" {
		serveJSON(jsonFile, mapHandler)
		return
	}

	if yamlFile != "" {
		serveYaml(yamlFile, mapHandler)
		return
	}

	fmt.Println("Using default map paths")
	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", mapHandler)
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func serveJSON(jsonFile string, fallback http.Handler) {
	jsonInput, err := os.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := handlers.JSONHandler(jsonInput, fallback)
	if err != nil {
		panic(err)
	}

	fmt.Println("Using json file paths")
	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", jsonHandler)
	if err != nil {
		panic(err)
	}
}

func serveYaml(yamlFile string, fallback http.Handler) {
	yaml, err := os.ReadFile(yamlFile)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := handlers.YAMLHandler(yaml, fallback)
	if err != nil {
		panic(err)
	}

	fmt.Println("Using yaml file paths")
	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", yamlHandler)
	if err != nil {
		panic(err)
	}
}
