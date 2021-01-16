package web

import (
	"net/http"
)

// BadRequest is used to respond back the status 400
func BadRequest(w http.ResponseWriter, err error) {
	var message []byte
	message, err = StructToBytes(err.Error)
	if err != nil {
		message = []byte("Unknown error occurred")
	}
	response(w, http.StatusBadRequest, message)
}

// Success returns the status code 200 with the response body
func Success(w http.ResponseWriter, body []byte) {
	response(w, http.StatusOK, body)
}

func response(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}
