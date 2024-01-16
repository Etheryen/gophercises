package handlers

import (
	"03-cyoa/types"
	"html/template"
	"net/http"
)

func init() {
	DEFAULT_TEMPLATE = template.Must(
		template.ParseFiles(DEFAULT_TEMPLATE_FILE),
	)
}

const DEFAULT_TEMPLATE_FILE string = "html/layout.html"

var DEFAULT_TEMPLATE *template.Template

func DEFAULT_PATH_FN(r *http.Request) string {
	path := r.URL.Path
	if path == "/" {
		path = "/intro"
	}
	return path[1:]
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(pathFn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = pathFn
	}
}

func NewHandler(
	storyMap map[string]types.StoryArc,
	opts ...HandlerOption,
) http.Handler {
	h := handler{storyMap, DEFAULT_TEMPLATE, DEFAULT_PATH_FN}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

type handler struct {
	storyMap map[string]types.StoryArc
	t        *template.Template
	pathFn   func(r *http.Request) string
}

func (h handler) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	path := h.pathFn(r)

	if story, ok := h.storyMap[path]; ok {
		// TODO: handle error
		_ = h.t.Execute(w, story)
		return
	}

	// TODO: handle 404
	http.Error(w, "Chapter not found", http.StatusNotFound)
}
