package propertymodel

import "github.com/ngleanhvu/go-booking/shared/core"

type PropertyCreateDTO struct {
	core.SQLModel
	Name           string  `form:"name" gorm:"column:name"`
	Star           float32 `form:"star_rating" gorm:"column:star_rating"`
	Address        string  `form:"address" gorm:"column:address"`
	CountryId      int     `form:"country_id" gorm:"column:country_id"`
	ProvinceId     int     `form:"province_id" gorm:"column:province_id"`
	WardId         int     `form:"ward_id" gorm:"column:ward_id"`
	Thumbnail      string  `form:"thumbnail" gorm:"column:thumbnail"`
	Lat            float64 `form:"lat" gorm:"column:lat"`
	Lng            float64 `form:"lng" gorm:"column:lng"`
	PropertyTypeId int     `form:"property_type_id" gorm:"column:property_type_id"`
}

func (PropertyCreateDTO) TableName() string {
	return "properties"
}

func (p *PropertyCreateDTO) Mask() {
	p.SQLModel.Mask(core.MaskTypeProperty)
}

type PropertyUpdateDTO struct {
	core.SQLModel
	Name           *string `form:"name" gorm:"column:name"`
	Address        *string `form:"address" gorm:"column:address"`
	ProvinceId     *int    `form:"province_id" gorm:"column:province_id"`
	WardId         *int    `form:"ward_id" gorm:"column:ward_id"`
	Thumbnail      string  `form:"thumbnail" gorm:"column:thumbnail"`
	PropertyTypeId *int    `form:"property_type_id" gorm:"column:property_type_id"`
}

func (PropertyUpdateDTO) TableName() string {
	return "properties"
}

func (p *PropertyUpdateDTO) Mask() {
	p.SQLModel.Mask(core.MaskTypeProperty)
}
