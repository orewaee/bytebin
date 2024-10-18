package middlewares

import (
	"github.com/rs/zerolog"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc, log *zerolog.Logger) http.Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		log.Debug().
			Str("method", request.Method).
			Str("path", request.URL.Path).
			Send()

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(handler)
}
