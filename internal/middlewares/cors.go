package middlewares

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(handler)
}
