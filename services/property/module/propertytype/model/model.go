package propertytypemodel

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyType struct {
	core.SQLModel
	Name string  `json:"name" form:"name" column:"name" db:"name"`
	Icon *string `json:"icon" form:"icon" column:"icon" db:"icon"`
}

func (PropertyType) TableName() string {
	return "property_types"
}

func (p *PropertyType) Mask() {
	p.SQLModel.Mask(core.MaskTypePropertyType)
}

func (p *PropertyType) Validate() error {
	name := strings.TrimSpace(p.Name)
	if err := checkPropertyName(name); err != nil {
		return err
	}
	return nil
}
