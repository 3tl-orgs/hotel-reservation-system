package model

import (
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type Facility struct {
	*core.SQLModel
	Icon *core.Image `json:"icon" gorm:"column:icon;type:jsonb"`
	Name string `json:"name" gorm:"column:name" db:"name"`
}

func (Facility) TableName() string {
	return "facilities"
}

func (f *Facility) Validate() error {
	f.Name = strings.TrimSpace(f.Name)

	if err := checkFacilityName(f.Name); err != nil {
		return err
	}

	return nil
}