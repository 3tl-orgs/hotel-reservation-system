package facilitypropertiesmodel

import (
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

type FacilityProperties struct {
	core.SQLModel
	PropertyId int                `json:"property_id" gorm:"column:property_id"`
	FacilityId int                `json:"facility_id" gorm:"column:facility_id"`
	AmenityIds core.JsonTypeArray `json:"amenity_ids" gorm:"column:amenity_ids"`
}

func (FacilityProperties) TableName() string {
	return "facility_properties"
}

func (fp *FacilityProperties) Mask() {
	fp.SQLModel.Mask(core.MaskTypeFacilityProperties)
}

type FacilityResponse struct {
	AmenityIds core.JsonTypeArray `json:"amenity_ids" gorm:"column:amenity_ids"`
	Status     bool               `json:"status" gorm:"column:status"`
	CreatedAt  time.Time          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time          `json:"updated_at" gorm:"column:updated_at"`
}
