package middleware

import (
	"context"
	"net/http"
	"strings"
)

func PathParams(next func(w http.ResponseWriter, r *http.Request), pattern string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxValues := getSegments(r.URL.Path, pattern)
		ctx := r.Context()
		for key, value := range ctxValues {
			ctx = context.WithValue(ctx, key, value)
		}
		next(w, r.WithContext(ctx))
	})
}

func getSegments(path string, pattern string) map[string]string {
	result := make(map[string]string)

	pathSegments := strings.Split(path, "/")
	patternSegments := strings.Split(pattern, "/")

	for idx, str := range patternSegments {
		if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
			key := str[1 : len(str)-1]
			result[key] = pathSegments[idx]
		}
	}

	return result
}
