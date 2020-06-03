package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	e, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(e)

	return nil
}

// ErrorResponse represents http error response structure required by datatable api
type errorResponse struct {
	error string // Optional: If an error occurs during the running of the server-side processing
}

// newErrorResponse creates a new error response to be consumed by the databale api
func newErrorResponse(err error) errorResponse {
	return errorResponse{error: err.Error()}
}

func sendError(w http.ResponseWriter, errMsg error) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Error: ", errMsg)
	response := newErrorResponse(errMsg)
	e, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(e)
}
