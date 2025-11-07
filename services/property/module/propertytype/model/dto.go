package model

import "strings"

type PropertyTypeCreateDto struct {
	Name string  `json:"name" db:"name"`
	Icon *string `json:"icon" db:"icon"`
}

func (PropertyTypeCreateDto) TableName() string {
	return PropertyType{}.TableName()
}

func (p *PropertyTypeCreateDto) Validate() error {
	name := strings.TrimSpace(p.Name)
	if err := checkPropertyName(name); err != nil {
		return err
	}
	return nil
}
