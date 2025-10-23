package model

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type CountryCreateDto struct {
	core.SQLModel
	Code string `json:"code" gorm:"column:code" db:"code"`
	Name string `json:"name" gorm:"column:name" db:"name"`
}

func (CountryCreateDto) TableName() string {
	return Country{}.TableName()
}

func (c *CountryCreateDto) Mask() {
	c.SQLModel.Mask(core.MaskTypeCountry)
}

func (c *CountryCreateDto) Validate() error {
	c.Code = strings.TrimSpace(c.Code)
	c.Name = strings.TrimSpace(c.Name)

	if err := checkCountryName(c.Name); err != nil {
		return ErrCountryNameIsEmpty
	}

	if err := checkCountryCode(c.Code); err != nil {
		return ErrCountryCodeIsEmpty
	}

	return nil
}

type CountryUpdateDto struct {
	Code *string `json:"code" gorm:"column:code" db:"code"`
	Name *string `json:"name" gorm:"column:name" db:"name"`
}

func (CountryUpdateDto) TableName() string {
	return Country{}.TableName()
}

func (c *CountryUpdateDto) Validate() error {
	*c.Code = strings.TrimSpace(*c.Code)
	*c.Name = strings.TrimSpace(*c.Name)

	if err := checkCountryName(*c.Name); err != nil {
		return ErrCountryNameIsEmpty
	}

	if err := checkCountryCode(*c.Code); err != nil {
		return ErrCountryCodeIsEmpty
	}

	return nil
}
