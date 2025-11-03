package core

import (
	"database/sql/driver"
	"encoding/json"
)

type JsonTypeArray []string

func (j JsonTypeArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonTypeArray) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), j)
}
