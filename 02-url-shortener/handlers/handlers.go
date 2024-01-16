package handlers

import (
	"02-url-shortener/utils"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(
	pathsToUrls map[string]string,
	fallback http.Handler,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, 308)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := utils.ParseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathMap := utils.BuildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(
	jsonInput []byte,
	fallback http.Handler,
) (http.HandlerFunc, error) {
	parsedJson, err := utils.ParseJSON(jsonInput)
	if err != nil {
		return nil, err
	}
	pathMap := utils.BuildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}
