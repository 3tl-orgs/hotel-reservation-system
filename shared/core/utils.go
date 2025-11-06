package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonTypeArray []string

func (j JsonTypeArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonTypeArray) Scan(src interface{}) error {
	if src == nil {
		*j = []string{}
		return nil
	}

	var data []byte
	switch v := src.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("unsupported type for JsonTypeArray: %T", src)
	}

	return json.Unmarshal(data, j)
}

// ToJSON: chuyển object bất kỳ thành JSON string
func ToJSON[T any](v T) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(data), nil
}

// MustToJSON: giống ToJSON nhưng panic khi lỗi
func MustToJSON[T any](v T) string {
	data, err := ToJSON(v)
	if err != nil {
		panic(err)
	}
	return data
}

// FromJSON: chuyển JSON string thành object generic
func FromJSON[T any](data string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return result, nil
}

// MustFromJSON: giống FromJSON nhưng panic khi lỗi
func MustFromJSON[T any](data string) T {
	result, err := FromJSON[T](data)
	if err != nil {
		panic(err)
	}
	return result
}

// ToMap: chuyển struct hoặc JSON string thành map[string]interface{}
func ToMap[T any](v T) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonStr, err := ToJSON(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to map: %w", err)
	}
	return result, nil
}

// ToArray: chuyển JSON string thành slice of generic
func ToArray[T any](data string) ([]T, error) {
	var result []T
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to array: %w", err)
	}
	return result, nil
}

// MustToMap: giống ToMap nhưng panic khi lỗi
func MustToMap[T any](v T) map[string]interface{} {
	result, err := ToMap(v)
	if err != nil {
		panic(err)
	}
	return result
}

// MustToArray: giống ToArray nhưng panic khi lỗi
func MustToArray[T any](data string) []T {
	result, err := ToArray[T](data)
	if err != nil {
		panic(err)
	}
	return result
}
