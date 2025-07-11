package handler

import (
	"encoding/json"
	"net/http"
)

type pathURL struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func MapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if dest, ok := urls[path]; ok {
			http.Redirect(writer, request, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	})
}

func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLS []pathURL
	err := json.Unmarshal(data, &pathURLS)
	if err != nil {
		return nil, err
	}

	urls := make(map[string]string)
	for _, pathURL := range pathURLS {
		urls[pathURL.Path] = pathURL.URL
	}

	return MapHandler(urls, fallback), nil
}
