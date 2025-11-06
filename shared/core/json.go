package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONType là generic type để lưu JSON vào DB qua GORM
type JSONType[T any] []T

// Value implement driver.Valuer
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

// Scan implement sql.Scanner
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
