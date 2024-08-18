package utils

import (
	"fmt"
	"net/http"
)

func MustWriteBytes(writer http.ResponseWriter, data []byte, code int) {
	writer.WriteHeader(code)

	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
}

func MustWriteString(writer http.ResponseWriter, data string, code int) {
	writer.WriteHeader(code)

	if _, err := fmt.Fprintln(writer, data); err != nil {
		panic(err)
	}
}
