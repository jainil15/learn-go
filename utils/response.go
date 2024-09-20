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
	if s.StatusCode == 0 {
		s.StatusCode = http.StatusOK
	}
	if s.Message == "" {
		s.Message = "Success"
	}
	if err != nil {
		return err
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(s.StatusCode)
	w.Write(val)
	return nil
}

func ErrorHandler(w http.ResponseWriter, e *ErrorResponse) {
	if e.StatusCode == 0 {
		e.StatusCode = http.StatusInternalServerError
	}
	// if e.Error == nil {
	// 	e.Error = map[string]interface{}{
	// 		"server": e.Message,
	// 	}
	// }
	val, err := json.Marshal(e)
	if err != nil {
		log.Fatalln(err)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(e.StatusCode)
	_, err = w.Write(val)
	if err != nil {
		return
	}
}
