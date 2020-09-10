package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResJSON(w http.ResponseWriter, statusCode int, res *Response) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if res.Message == "" {
		res.Message = http.StatusText(statusCode)
	}
	return json.NewEncoder(w).Encode(res)
}

func ResERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ResJSON(w, statusCode, &Response{
			Message: err.Error(),
		})
		return
	}
	ResJSON(w, http.StatusBadRequest, &Response{})
}
