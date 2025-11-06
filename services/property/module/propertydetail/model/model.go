package propertydetailmodel

import "github.com/ngleanhvu/go-booking/shared/core"

type PropertyDetailModel struct {
	core.SQLModel
	PropertyTypeId int     `json:"property_type_id" gorm:"column:property_type_id"`
	CheckInTime    *string `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime   *string `json:"check_out_time" gorm:"column:check_out_time"`
	Description    *string `json:"description" gorm:"column:description"`
	Hotline        *string `json:"hotline" gorm:"column:hotline"`
	HostId         *int    `json:"host_id" gorm:"column:host_id"`
	WebsiteUrl     *string `json:"website_url" gorm:"column:website_url"`
	Email          *string `json:"email" gorm:"column:email"`
	Images         *string `json:"images" gorm:"column:images"`
	BestAmenities  *string `json:"best_amenities" gorm:"column:best_amenities"`
}

func (PropertyDetailModel) TableName() string {
	return "property_details"
}
