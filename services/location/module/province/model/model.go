package provincemodel

import "github.com/ngleanhvu/go-booking/shared/core"

type Province struct {
	core.SQLModel
	Name      string `json:"name" gorm:"column:name"`
	Code      string `json:"code" gorm:"column:code;unique"`
	CountryId int    `json:"country_id" gorm:"column:country_id"`
}

func (Province) TableName() string {
	return "provinces"
}

func (p *Province) Mask() {
	p.SQLModel.Mask(core.MaskTypeProvince)
}
