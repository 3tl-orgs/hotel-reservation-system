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
