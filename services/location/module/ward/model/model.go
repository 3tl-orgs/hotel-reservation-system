package wardsmodel

import "github.com/ngleanhvu/go-booking/shared/core"

type Ward struct {
	core.SQLModel
	Name       string `json:"name" gorm:"column:name"`
	Code       string `json:"code" gorm:"column:code;unique"`
	ProvinceId int    `json:"province_id" gorm:"column:province_id"`
}

func (Ward) TableName() string {
	return "wards"
}

func (ward *Ward) Mask() {
	ward.SQLModel.Mask(core.MaskTypeWard)
}
