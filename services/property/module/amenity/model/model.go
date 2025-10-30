package model

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type Amenity struct {
	*core.SQLModel
	icon *string `json:"icon" gorm:"column:icon" db:"icon"`
	Name string  `json:"name" gorm:"column:name" db:"name"`
}

func (Amenity) TableName() string {
	return "amenities"
}

func (a *Amenity) Validate() error {
	a.Name = strings.TrimSpace(a.Name)

	if err := checkAmenityName(a.Name); err != nil {
		return err
	}

	return nil
}
