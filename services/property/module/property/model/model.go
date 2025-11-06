package propertymodel

import "github.com/ngleanhvu/go-booking/shared/core"

type Property struct {
	core.SQLModel
	Name string  `json:"name" gorm:"column:name"`
	Star float32 `json:"star_rating" gorm:"column:star_rating"`
	
	Address        string  `json:"address" gorm:"column:address"`
	CountryId      int     `json:"country_id" gorm:"column:country_id"`
	ProvinceId     int     `json:"province_id" gorm:"column:province_id"`
	WardId         int     `json:"ward_id" gorm:"column:ward_id"`
	Thumbnail      string  `json:"thumbnail" gorm:"column:thumbnail"`
	Lat            float64 `json:"lat" gorm:"column:lat"`
	Lng            float64 `json:"lng" gorm:"column:lng"`
	PropertyTypeId int     `json:"property_type_id" gorm:"column:property_type_id"`
}

func (Property) TableName() string {
	return "properties"
}

func (p *Property) Mask() {
	p.SQLModel.Mask(core.MaskTypeProperty)
}
