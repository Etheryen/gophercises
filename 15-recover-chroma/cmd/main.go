package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n")

	for i, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}

		words := strings.Split(line, " ")
		info := strings.Split(strings.TrimSpace(words[0]), ":")
		file := info[0]
		lineNum := info[1]

		v := url.Values{}
		v.Set("file", file)
		v.Set("line", lineNum)

		words[0] = fmt.Sprintf(
			"\t<a href=\"/debug/?%s\">%s:%s</a>",
			v.Encode(),
			file,
			lineNum,
		)

		lines[i] = strings.Join(words, " ")
	}

	return strings.Join(lines, "\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", sourceCodeHandler)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("file")
	lineStr := r.URL.Query().Get("line")

	line, err := strconv.Atoi(lineStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lexer := lexers.Match(path)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get("nord")
	if style == nil {
		style = styles.Fallback
	}

	formatter := html.New(
		html.Standalone(true),
		html.WithLineNumbers(true),
		html.LineNumbersInTable(true),
		html.TabWidth(4),
		html.HighlightLines([][2]int{{line, line}}),
	)

	iterator, err := lexer.Tokenise(nil, string(file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := formatter.Format(w, style, iterator); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				log.Println(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)

				fmt.Fprintf(
					w,
					"<h1>panic: %v</h1><pre>%s</pre>",
					err,
					makeLinks(string(debug.Stack())),
				)
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
