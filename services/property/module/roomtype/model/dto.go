package roomtypemodel

import "github.com/ngleanhvu/go-booking/shared/core"

type RoomTypeCreateDTO struct {
	core.SQLModel
	PropertyId    int                `json:"property_id" gorm:"column:property_id"`
	Name          string             `json:"name" gorm:"column:name"`
	Description   string             `json:"description" gorm:"column:description"`
	Images        core.JsonTypeArray `json:"images" gorm:"column:images"`
	TotalRooms    int                `json:"total_rooms" gorm:"column:total_rooms"`
	MaxAdults     int                `json:"max_adults" gorm:"column:max_adults"`
	MaxChildren   int                `json:"max_children" gorm:"column:max_children"`
	BasePrice     float64            `json:"base_price" gorm:"column:base_price"`
	PayBefore     bool               `json:"pay_before" gorm:"column:pay_before"`
	BedType       string             `json:"bed_type" gorm:"column:bed_type"`
	PayType       string             `json:"pay_type" gorm:"column:pay_type"`
	BestAmenities core.JsonTypeArray `json:"best_amenities" gorm:"column:best_amenities"`
	Amenities     core.JsonTypeArray `json:"amenities" gorm:"column:amenities"`
}

func (RoomTypeCreateDTO) TableName() string {
	return "room_types"
}

func (r *RoomTypeCreateDTO) Mask() {
	r.SQLModel.Mask(core.MaskTypeRoomType)
}

type RoomTypeUpdateDTO struct {
	core.SQLModel
	PropertyId    *int                `json:"property_id" gorm:"column:property_id"`
	Name          *string             `json:"name" gorm:"column:name"`
	Description   *string             `json:"description" gorm:"column:description"`
	Images        *core.JsonTypeArray `json:"images" gorm:"column:images"`
	TotalRooms    *int                `json:"total_rooms" gorm:"column:total_rooms"`
	MaxAdults     *int                `json:"max_adults" gorm:"column:max_adults"`
	MaxChildren   *int                `json:"max_children" gorm:"column:max_children"`
	BasePrice     *float64            `json:"base_price" gorm:"column:base_price"`
	PayBefore     *bool               `json:"pay_before" gorm:"column:pay_before"`
	BedType       *string             `json:"bed_type" gorm:"column:bed_type"`
	PayType       *string             `json:"pay_type" gorm:"column:pay_type"`
	BestAmenities *core.JsonTypeArray `json:"best_amenities" gorm:"column:best_amenities"`
	Amenities     *core.JsonTypeArray `json:"amenities" gorm:"column:amenities"`
}

func (RoomTypeUpdateDTO) TableName() string {
	return "room_types"
}

func (r *RoomTypeUpdateDTO) Mask() {
	r.SQLModel.Mask(core.MaskTypeRoomType)
}

func (r *RoomTypeUpdateDTO) PatchToModel() map[string]interface{} {
	updated := make(map[string]interface{})
	if r.Name != nil {
		updated["name"] = *r.Name
	}

	if r.Description != nil {
		updated["description"] = *r.Description
	}

	if r.Images != nil {
		updated["images"] = *r.Images
	}

	if r.TotalRooms != nil {
		updated["total_rooms"] = *r.TotalRooms
	}

	if r.MaxAdults != nil {
		updated["max_adults"] = *r.MaxAdults
	}

	if r.MaxChildren != nil {
		updated["max_children"] = *r.MaxChildren
	}

	if r.BasePrice != nil {
		updated["base_price"] = *r.BasePrice
	}

	if r.PayBefore != nil {
		updated["pay_before"] = *r.PayBefore
	}

	if r.BedType != nil {
		updated["bed_type"] = *r.BedType
	}

	if r.PayType != nil {
		updated["pay_type"] = *r.PayType
	}

	if r.BestAmenities != nil {
		updated["best_amenities"] = *r.BestAmenities
	}

	if r.Amenities != nil {
		updated["amenities"] = *r.Amenities
	}

	return updated
}
