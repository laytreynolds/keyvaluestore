package store

type Storer interface {
	Get(key string, value []byte) (any, error)
	Add()
}

type KVStore struct {
	store map[string]any
}

type Response struct {
	Value any
	Error error
}

func NewKeyValueStore() *KVStore {
	return &KVStore{
		store: make(map[string]any), // Initialising the map with make
	}
}

func (s *KVStore) InitData() {

	// slice[i] :: ptr + i * size_of_data
	m := map[string]any{
		"TestString": []byte(`"Value1"`),                      // Directly using byte slices for strings
		"TestNumber": []byte(`1`),                             // Directly using byte slices for numbers
		"TestMap":    []byte(`{"name": "layton", "age": 27}`), // JSON string for a map
	}

	for k, v := range m {
		// Add the byte slice value to the store
		s.Add(k, v.([]byte))
	}
}
