package model

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type Country struct {
	*core.SQLModel `json:",inline"`
	Name           string `json:"name" gorm:"column:name"`
	Code           string `json:"code" gorm:"column:code,unique"`
}

func (Country) TableName() string {
	return "countries"
}

func (c *Country) Mask() {
	c.SQLModel.Mask(core.MaskTypeCountry)
}

func (c *Country) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	c.Code = strings.TrimSpace(c.Code)

	if err := checkCountryName(c.Name); err != nil {
		return err
	}

	if err := checkCountryCode(c.Code); err != nil {
		return err
	}

	return nil
}
