package store

import (
	"errors"
	"kvstore/helpers"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		description string
		key         string
		value       []byte
		want        any
	}{
		{
			description: "Test String",
			key:         "AddString",
			value:       []byte(`"Layton"`),
			want:        "Layton",
		},
		{
			description: "Test Number",
			key:         "AddNumber",
			value:       []byte(`1`),
			want:        float64(1),
		},
		{
			description: "Test Map",
			key:         "TestMap",
			value:       []byte(`{"name": "Layton"}`),
			want:        map[string]any{"name": "Layton"},
		},
		{
			description: "TestDuplicate",
			key:         "AddString",
			value:       []byte(`{"name": "Layton"}`),
			want:        helpers.DuplicateKeyError,
		},
		{
			description: "TestParseFail",
			key:         "ParseFail",
			value:       []byte(`{"name": "Layton}`),
			want:        errors.New("unexpected end of JSON input"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := Store.Add(tt.key, tt.value)

			// Type assertion for the expected type
			switch expected := tt.want.(type) {
			case string:
				if gotStr, ok := got.(string); !ok || gotStr != expected {
					t.Errorf("Add() = %v, want %v", got, expected)
				}
			case float64:
				if gotFloat, ok := got.(float64); !ok || gotFloat != expected {
					t.Errorf("Add() = %v, want %v", got, expected)
				}
			case map[string]any:
				if gotMap, ok := got.(map[string]any); !ok || !compareMaps(gotMap, expected) {
					t.Errorf("Add() = %v, want %v", got, expected)
				}
			case error:
				if err == nil {
					t.Errorf("Add() = %v, want %v", err, expected)
				}
			default:
				t.Errorf("Unsupported type for comparison: %T", expected)
			}
		})
	}
}

func TestGet(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		want        any
	}{
		{
			description: "TestString",
			key:         "TestString",
			want:        "Value1",
		},
		{
			description: "TestNumber",
			key:         "TestNumber",
			want:        float64(1),
		},
		{
			description: "TestMap",
			key:         "TestMap",
			want:        map[string]any{"name": "layton", "age": float64(27)},
		},
		{
			description: "TestNotExist",
			key:         "NotExist",
			want:        helpers.NotExistError,
		},
	}

	// Test getting an existing key
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := store.Get(tt.key)

			// Type assertion for the expected type
			switch expected := tt.want.(type) {
			case string:
				if gotStr, ok := got.(string); !ok || gotStr != expected {
					t.Errorf("Get() = %v, want %v", got, expected)
				}
			case float64:
				if gotFloat, ok := got.(float64); !ok || gotFloat != expected {
					t.Errorf("Get() = %v, want %v", got, expected)
				}
			case map[string]any:
				if gotMap, ok := got.(map[string]any); !ok || !compareMaps(gotMap, expected) {
					t.Errorf("Get() = %v, want %v", got, expected)
				}
			case error:
				if err == nil {
					t.Errorf("Update() error = %v, want %v", err, expected)
				}
			default:
				t.Errorf("Unsupported type for comparison: %T", expected)
			}
		})
	}
}

func TestExists(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		want        any
	}{
		{
			description: "TestString",
			key:         "TestString",
			want:        true,
		},
		{
			description: "TestNumber",
			key:         "TestNumber",
			want:        true,
		},
		{
			description: "TestMap",
			key:         "TestMap",
			want:        true,
		},
		{
			description: "TestFalse",
			key:         "TestFalse",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, _ := store.Exists(tt.key)

			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}

	// Check we receive the correct error
	want := helpers.NotExistError
	key := "TestFalse"
	if _, err := store.Exists(key); err != want {
		t.Errorf("Exists() = %v, want %v", err, want)
	}
}

func TestCount(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		want        any
	}{
		{
			description: "TestCount",
			want:        len(store.store),
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, _ := store.Count()
			if got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		want        any
	}{
		{
			description: "TestDelete",
			key:         "TestString",
			want:        nil,
		},
		{
			description: "TestDeleteNotExist",
			key:         "NotExist",
			want:        helpers.NotExistError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := store.Delete(tt.key)

			if got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		value       []byte
		want        any
	}{
		{
			description: "Test String",
			key:         "TestString",
			value:       []byte(`"Layton"`),
			want:        "Layton",
		},
		{
			description: "Test Number",
			key:         "TestNumber",
			value:       []byte(`2`),
			want:        float64(2),
		},
		{
			description: "Test Map",
			key:         "TestMap",
			value:       []byte(`{"name": "Layton"}`),
			want:        map[string]any{"name": "Layton"},
		},
		{
			description: "TestNotExist",
			key:         "NotExist",
			value:       []byte(`value`),
			want:        helpers.NotExistError,
		},
		{
			description: "TestParseFail",
			key:         "ParseFail",
			value:       []byte(`{"name": "Layton}`),
			want:        errors.New("unexpected end of JSON input"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := store.Update(tt.key, tt.value)

			switch expected := tt.want.(type) {
			case string:
				if gotStr, ok := got.(string); !ok || gotStr != expected {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			case float64:
				if gotFloat, ok := got.(float64); !ok || gotFloat != expected {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			case map[string]any:
				if gotMap, ok := got.(map[string]any); !ok || !compareMaps(gotMap, expected) {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			case error:
				if err == nil {
					t.Errorf("Update() error = %v, want %v", err, expected)
				}
			}
		})
	}
}

func TestClear(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		want        map[string]any
	}{
		{
			description: "TestClear",
			want:        make(map[string]any),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := store.Clear()
			if err != nil {
				t.Errorf("Clear() returned an error: %v", err)
			}

			// Type assertion for 'got' to map[string]any
			gotMap, ok := got.(map[string]any)
			if !ok {
				t.Errorf("Clear() did not return a map[string]any, got %T", got)
				return
			}

			if !compareMaps(gotMap, tt.want) {
				t.Errorf("Clear() = %v, want %v", gotMap, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		want        int
	}{
		{
			description: "TestGetAll",
			want:        len(store.store),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := store.GetAll()
			if err != nil {
				t.Errorf("GetAll() returned an error: %v", err)
			}

			// Type assertion to convert any to map[string]any
			gotMap, ok := got.(map[string]any)
			if !ok {
				t.Errorf("GetAll() did not return a map[string]any, got %T", got)
				return
			}

			count := len(gotMap) // Count the number of items in the map
			if count != tt.want {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpsert(t *testing.T) {
	store := NewKeyValueStore()
	store.InitData()

	tests := []struct {
		description string
		key         string
		value       []byte
		want        any
	}{
		{
			description: "Test String Update",
			key:         "TestString",
			value:       []byte(`"James"`),
			want:        "James",
		},
		{
			description: "Test Number Update",
			key:         "TestNumber",
			value:       []byte(`89`),
			want:        float64(89),
		},
		{
			description: "Test Map Update",
			key:         "TestMap",
			value:       []byte(`{"name": "Jimmy"}`),
			want:        map[string]any{"name": "Jimmy"},
		},
		{
			description: "TestUpsert",
			key:         "NewKey",
			value:       []byte(`"NewValue"`),
			want:        "NewValue",
		},
		{
			description: "TestParseFail",
			key:         "ParseFail",
			value:       []byte(`{"name": "Layton}`),
			want:        errors.New("unexpected end of JSON input"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, _ := store.Upsert(tt.key, tt.value)

			switch expected := tt.want.(type) {
			case string:
				if gotStr, ok := got.(string); !ok || gotStr != expected {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			case float64:
				if gotFloat, ok := got.(float64); !ok || gotFloat != expected {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			case map[string]any:
				if gotMap, ok := got.(map[string]any); !ok || !compareMaps(gotMap, expected) {
					t.Errorf("Update() = %v, want %v", got, expected)
				}
			}
		})
	}
}

func compareMaps(got, want map[string]any) bool {
	if len(got) != len(want) {
		return false
	}
	for key, wantValue := range want {
		gotValue, ok := got[key]
		if !ok || !equal(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// equal checks if two any values are equal
func equal(a, b any) bool {
	return a == b
}
