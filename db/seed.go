package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	baseURL = "http://localhost:8080/kvs/add" // Change this to your server's URL
)

var start = time.Now()

// Seed generates a specified number of random key-value pairs to populate the database.
func Seed(pairs int) {
	// Create a new random generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var count int

	for count = 0; count < pairs; count++ {
		key := fmt.Sprintf("key%d", count+1)
		value := fmt.Sprintf("value%d", rng.Intn(20000)) // Random value between 0 and 999
		_ = addKeyValue(key, value)
	}

	elapsed := time.Since(start)
	fmt.Printf("%d keys added. Time Taken: %s", count, elapsed)

}

// addKeyValue sends a POST request to the /add endpoint with the specified key and value.
func addKeyValue(key, value string) error {
	// Create the request body
	body := map[string]string{"value": value}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", baseURL, key), bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	return nil
}
