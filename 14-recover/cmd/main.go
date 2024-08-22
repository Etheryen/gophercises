package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
)

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New(
			"the ResponseWriter does not implement the Hijacker interface",
		)
	}
	return hijacker.Hijack()
}

func (rw *responseWriter) Flush() {
	if flusher, ok := rw.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (rw *responseWriter) flush() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}

	for _, write := range rw.writes {
		if _, err := rw.ResponseWriter.Write(write); err != nil {
			return err
		}
	}

	return nil
}

func recoverMw(handler http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
				log.Println("Stack:", string(debug.Stack()))
				if !dev {
					http.Error(
						w,
						"Something went wrong",
						http.StatusInternalServerError,
					)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(
					w,
					"<h1>panic: %v</h1><pre>%s</pre>",
					err,
					string(debug.Stack()),
				)
			}
		}()

		nw := &responseWriter{ResponseWriter: w}
		handler.ServeHTTP(nw, r)
		nw.flush()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", recoverMw(mux, true)))
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
