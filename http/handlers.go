package http

import (
	"encoding/json"
	"fmt"
	"io"
	"kvstore/channels"
	"kvstore/helpers"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serviceName = "kvstore"
	version     = "0"
)

// Ping returns service name, version and hostname of service.
func Ping(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodGet); err != nil {
		helpers.HandleError(w, err)
		return
	}

	hostname, _ := os.Hostname()

	res := struct {
		ServiceName string `json:"service_name"`
		Version     string `json:"version"`
		HostName    string `json:"hostname"`
		DateTime    string `json:"datetime"`
	}{
		ServiceName: serviceName,
		Version:     version,
		HostName:    hostname,
		DateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Get processes the incoming HTTP request.
func Get(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodGet); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, _, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.GetRequest(k)
	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}

	log.Printf("Successfully Received Key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)
}

// Add Calls store.Add to Add a key to the map
func Add(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodPost); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, v, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.AddRequest(k, v)

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}

	log.Printf("Successfully Added value under key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)

}

// GetAll calls store.GetAll and returns the full store
func GetAll(w http.ResponseWriter, r *http.Request) {
	err := helpers.CheckMethod(r.Method, http.MethodGet)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}
	resp := channels.GetAllRequest()

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully retrieved All keys")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)
}

// Exists checks membership in the store for a key
func Exists(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodGet); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, _, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.ExistsRequest(k)

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully checked membership for key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)

}

// Count calls store.Count and returns the count of objects in the store
func Count(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodGet); err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.CountRequest()

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully Counted All keys")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)
}

// Clear calls store.Clear clears the store
func Clear(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodPost); err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.ClearRequest()

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully Cleared All keys")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Store Cleared\n"))
}

// Delete calls store.Delete deletes the item from the store
func Delete(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodDelete); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, _, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.DeleteRequest(k)

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully deleted key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Key Deleted\n"))
}

// Update calls store.Update and updates a value from the store
func Update(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodPut); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, v, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.UpdateRequest(k, v)

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully updated key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)
}

// Upsert calls store.Upsert and updates a value from the store or inserts if it does not exist.
func Upsert(w http.ResponseWriter, r *http.Request) {
	if err := helpers.CheckMethod(r.Method, http.MethodPut); err != nil {
		helpers.HandleError(w, err)
		return
	}

	k, v, err := GetParam(r)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	resp := channels.UpsertRequest(k, v)

	if resp.Error != nil {
		helpers.HandleError(w, resp.Error)
		return
	}
	log.Printf("Successfully Updated key: %s", k)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Value)

}

// GetParam takes in an HTTP request and returns a key, value, and error (maybe nil).
func GetParam(r *http.Request) (string, []byte, error) {

	var value []byte
	var err error

	defer func(Body io.ReadCloser) {
		Body.Close()
	}(r.Body)

	switch r.Method {
	case http.MethodPost, http.MethodPut:
		// Read values from the body
		value, err = io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error decoding body: %v", err)
			return "", nil, fmt.Errorf("failed to decode body: %w", err)
		}

		if len(value) == 0 {
			return "", nil, helpers.MissingValueError
		}

	default:
		value = nil
	}

	/// Parse the URL form for both GET and POST requests
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v", err)
		return "", nil, fmt.Errorf("failed to parse form: %w", err)
	}

	// GetRequest Key from URL request
	key := r.Form.Get("key")
	if key == "" {
		return "", nil, helpers.MissingKeyError
	}

	// Return the key, value, and nil error
	return key, value, nil
}
