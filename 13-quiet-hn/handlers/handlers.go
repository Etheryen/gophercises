package handlers

import (
	"13-quiet-hn/cache"
	"fmt"
	"html/template"
	"net/http"
)

func GetRoot(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html"))

	data, err := cache.GetTemplateData()
	if err != nil {
		handleErr(w, err)
		return
	}

	if err = t.Execute(w, data); err != nil {
		handleErr(w, err)
	}
}

func handleErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal server error:", err)
}
