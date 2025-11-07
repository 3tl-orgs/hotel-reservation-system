package model

import (
	"strings"
	"time"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type FacilityCreateDto struct {
	*core.SQLModel
	Icon *core.Image `json:"icon" gorm:"column:icon;type:jsonb"`
	Name string      `form:"name" json:"name"`
}

func (FacilityCreateDto) TableName() string {
	return Facility{}.TableName()
}

func (f *FacilityCreateDto) Validate() error {
	f.Name = strings.TrimSpace(f.Name)

	if err := checkFacilityName(f.Name); err != nil {
		return err
	}
	return nil
}

type FacilityUpdateDto struct {
	Icon      *core.Image `json:"icon" gorm:"column:icon" db:"icon"`
	Name      *string     `json:"name" gorm:"column:name" db:"name"`
	UpdatedAt *time.Time  `json:"updated_at" gorm: "column:updated_at" db:"updated_at"`
}

func (FacilityUpdateDto) TableName() string { //Xác định update cho table nào
	return Facility{}.TableName()
}

func (f *FacilityUpdateDto) Validate() error {
	if f.Name != nil {
		*f.Name = strings.TrimSpace(*f.Name) //Thay đổi dữ liêu trong ô bị trỏ

		if err := checkFacilityName(*f.Name); err != nil {
			return err
		}
	}
	return nil
}
