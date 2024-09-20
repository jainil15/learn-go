package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Result     interface{} `json:"result"`
	Message    string      `json:"message"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
	Message    string      `json:"message"`
}

func ResponseHandler(w http.ResponseWriter, s *SuccessResponse) error {
	val, err := json.Marshal(s)
	if err != nil {
		return err
	}
	w.WriteHeader(s.StatusCode)
	w.Header().Add("Content-type", "application/json")
	w.Write(val)
	return nil
}

func ErrorHandler(w http.ResponseWriter, e *ErrorResponse) {
	val, err := json.Marshal(e)
	if err != nil {
		log.Fatalln(err)
		return
	}
	w.WriteHeader(e.StatusCode)
	w.Header().Add("Content-type", "application/json")
	_, err = w.Write(val)
	if err != nil {
		return
	}
}
