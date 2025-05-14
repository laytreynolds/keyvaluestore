package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Error types
var (
	MissingKeyError   = errors.New("key not provided")
	EmptyValueError   = errors.New("value not found")
	MissingValueError = errors.New("value not provided")
	NotExistError     = errors.New("key not found")
	DuplicateKeyError = errors.New("duplicate key")
	MethodNotAllowed  = errors.New("method not allowed")
)

// ParseJSON takes in a byte array and parses into an any
func ParseJSON(body []byte) (any, error) {
	var value any

	err := json.Unmarshal(body, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// HandleError takes in an error, logs out the error and returns a status code
func HandleError(w http.ResponseWriter, err error) {
	if errors.Is(err, MissingKeyError) {
		log.Printf("Key Error: %s", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if errors.Is(err, MissingValueError) {
		log.Printf("Value Error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, NotExistError) {
		log.Printf("Key Error: %s", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if errors.Is(err, DuplicateKeyError) {
		log.Printf("Key Error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, EmptyValueError) {
		log.Printf("Value Error: %s", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if errors.Is(err, MethodNotAllowed) {
		log.Printf("Method Error: %s", err)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Handle other errors if necessary
	log.Printf("Unexpected Error: %s", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// CheckMethod checks the method against a valid "want" value
func CheckMethod(method, want string) error {
	if method != want {
		return MethodNotAllowed
	}
	return nil
}
