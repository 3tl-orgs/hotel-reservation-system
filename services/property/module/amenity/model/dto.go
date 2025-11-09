package amenitymodel

import (
	"strings"
	"time"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type AmenityCreateDto struct {
	*core.SQLModel
	Icon *string `json:"icon" gorm:"icon" db:"icon"`
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

	if err := checkAmenityName(a.Name); err != nil {
		return err
	}

	return nil
}

type AmenityUpdateDto struct {
	Icon      *string    `json:"icon" gorm:"column:icon" db:"icon"`
	Name      *string    `json:"name" gorm:"column:name" db:"name"`
	UpdatedAt *time.Time `json:"updated_at" gorm: "column:updated_at" db:"updated_at"`
}

func (AmenityUpdateDto) TableName() string { //Xác định update cho table nào
	return Amenity{}.TableName()
}

func (a *AmenityUpdateDto) Validate() error {
	if a.Name != nil {
		*a.Name = strings.TrimSpace(*a.Name)

		if err := checkAmenityName(*a.Name); err != nil {
			return err
		}
	}
	return nil
}
