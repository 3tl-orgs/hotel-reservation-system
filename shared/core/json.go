package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// ✅ JSONType: lưu mảng JSON (ví dụ: ["A", "B"])
type JSONType[T any] []T

func (j JSONType[T]) Value() (driver.Value, error) {
	if j == nil {
		return "[]", nil
	}
	data, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSONType: %w", err)
	}
	return string(data), nil
}

func (j *JSONType[T]) Scan(src interface{}) error {
	if src == nil {
		*j = []T{}
		return nil
	}

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("unsupported type for JSONType: %T", src)
	}

	return json.Unmarshal(data, j)
}

// ✅ JSONObject: lưu một object JSON duy nhất (ví dụ: {"name":"Vũ","age":20})
type JSONObject[T any] struct {
	Data *T
}

// Khi ghi xuống DB
func (j JSONObject[T]) Value() (driver.Value, error) {
	if j.Data == nil {
		return "{}", nil
	}
	data, err := json.Marshal(j.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSONObject: %w", err)
	}
	return string(data), nil
}

// Khi đọc từ DB
func (j *JSONObject[T]) Scan(src interface{}) error {
	if src == nil {
		j.Data = new(T)
		return nil
	}

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("unsupported type for JSONObject: %T", src)
	}

	var obj T
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("failed to unmarshal JSONObject: %w", err)
	}
	j.Data = &obj
	return nil
}
