package facilitypropertiesmodel

import "github.com/ngleanhvu/go-booking/shared/core"

type FacilityPropertiesCreateDTO struct {
	core.SQLModel
	PropertyId int                `json:"property_id" gorm:"column:property_id"`
	FacilityId int                `json:"facility_id" gorm:"column:facility_id"`
	AmenityIds core.JsonTypeArray `json:"amenity_ids" gorm:"column:amenity_ids"`
}

func (FacilityPropertiesCreateDTO) TableName() string {
	return "facility_properties"
}

func (fp *FacilityPropertiesCreateDTO) Mask() {
	fp.SQLModel.Mask(core.MaskTypeFacilityProperties)
}

func (fp *FacilityPropertiesCreateDTO) Validate() error {
	if fp.FacilityId <= 0 || fp.PropertyId <= 0 {
		return ErrIdOrAmenitiesNotAccepted
	}

	return nil
}

type FacilityPropertiesUpdateDTO struct {
	core.SQLModel
	Status     *bool               `json:"status" gorm:"column:status"`
	AmenityIds *core.JsonTypeArray `json:"amenity_ids" gorm:"column:amenity_ids"`
}

func (FacilityPropertiesUpdateDTO) TableName() string {
	return "facility_properties"
}

func (fp *FacilityPropertiesUpdateDTO) Mask() {
	fp.SQLModel.Mask(core.MaskTypeFacilityProperties)
}
