package propertydetailmodel

import "github.com/ngleanhvu/go-booking/shared/core"

type PropertyDetail struct {
	core.SQLModel
	PropertyId    int                        `json:"property_id" gorm:"column:property_id"`
	CheckInTime   *string                    `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime  *string                    `json:"check_out_time" gorm:"column:check_out_time"`
	Description   *string                    `json:"description" gorm:"column:description"`
	Hotline       *string                    `json:"hotline" gorm:"column:hotline"`
	HostId        *int                       `json:"host_id" gorm:"column:host_id"`
	WebsiteUrl    *string                    `json:"website_url" gorm:"column:website_url"`
	Email         *string                    `json:"email" gorm:"column:email"`
	Images        *core.JSONType[core.Image] `json:"images" gorm:"column:images"`
	BestAmenities *core.JSONType[int]        `json:"best_amenities" gorm:"column:best_amenities"`
}

func (PropertyDetail) TableName() string {
	return "property_details"
}
