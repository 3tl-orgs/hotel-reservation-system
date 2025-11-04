package core

import (
	"database/sql/driver"
	"encoding/json"
)

// JsonTypeArray Tự động cast về jsonb
type JsonTypeArray []string

// Value implement lại hàm built-in
func (j JsonTypeArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan implement lại hàm built-in
func (j *JsonTypeArray) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), j)
}
