// Package store provides core key/value store functionality. The store provides a map of [string]any
package store

import (
	"kvstore/helpers"
)

var (
	Store = NewKeyValueStore()
)

func (s *KVStore) Get(key string) (any, error) {

	value, ok := s.store[key]
	if !ok {
		return "", helpers.NotExistError
	}
	return value, nil
}

func (s *KVStore) Add(key string, v []byte) (any, error) {

	// Check for duplicate keys
	if _, ok := s.store[key]; ok {
		return "", helpers.DuplicateKeyError // Return early if the key already exists
	}

	// Parse the value from JSON
	value, err := helpers.ParseJSON(v)
	if err != nil {
		return "", err // Return early if parsing fails
	}
	// Add the key-value pair to the store
	s.store[key] = value

	return value, nil
}

func (s *KVStore) GetAll() (any, error) {

	return s.store, nil
}

func (s *KVStore) Exists(key string) (bool, error) {

	if _, ok := s.store[key]; !ok {
		return false, helpers.NotExistError
	}

	return true, nil
}

func (s *KVStore) Count() (int, error) {

	return len(s.store), nil
}

func (s *KVStore) Clear() (any, error) {

	clear(s.store)

	return s.store, nil
}

func (s *KVStore) Delete(key string) error {

	if _, ok := s.store[key]; !ok {
		return helpers.NotExistError

	}

	delete(s.store, key)
	return nil
}

func (s *KVStore) Update(key string, v []byte) (any, error) {

	if _, ok := s.store[key]; !ok {
		return "", helpers.NotExistError
	}

	// Parse the value from JSON
	value, err := helpers.ParseJSON(v)
	if err != nil {
		return "", err // Return early if parsing fails
	}

	s.store[key] = value

	return value, nil
}

func (s *KVStore) Upsert(key string, v []byte) (any, error) {

	// Parse the value from JSON
	value, err := helpers.ParseJSON(v)
	if err != nil {
		return "", err // Return early if parsing fails
	}

	s.store[key] = value

	return value, nil
}
