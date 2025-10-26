package model

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type AmenityCreateDto struct {
	*core.SQLModel
	icon *string `json:"icon" gorm:"icon" db:"icon"`
	Name string  `json:"name" gorm:"name" db:"name"`
}

func (AmenityCreateDto) TableName() string {
	return Amenity{}.TableName()
}

func (a *AmenityCreateDto) Mask() {
	a.SQLModel.Mask(core.MaskTypeAmenity)
}

func (a *AmenityCreateDto) Validate() error {
	a.Name = strings.TrimSpace(a.Name)

	if err := checkName(a.Name); err != nil {
		return err
	}

	return nil
}
