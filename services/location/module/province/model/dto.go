package provincemodel

import (
	"github.com/ngleanhvu/go-booking/shared/core"
)

type CreateProvinceDTO struct {
	core.SQLModel
	Name      string `json:"name" gorm:"column:name"`
	Code      string `json:"code" gorm:"column:code"`
	CountryId int    `json:"country_id" gorm:"column:country_id"`
}

func (CreateProvinceDTO) TableName() string {
	return "provinces"
}

func (p *CreateProvinceDTO) Mask() {
	p.SQLModel.Mask(core.MaskTypeProvince)
}

type UpdateProvinceDTO struct {
	core.SQLModel
	Name      string `json:"name" gorm:"column:name"`
	Code      string `json:"code" gorm:"column:code"`
	CountryId int    `json:"country_id" gorm:"column:country_id"`
}

func (UpdateProvinceDTO) TableName() string {
	return "provinces"
}

func (p *UpdateProvinceDTO) Mask() {
	p.SQLModel.Mask(core.MaskTypeProvince)
}

//func (p *UpdateProvinceDTO) Validate() error {
//	p.Code = strings.TrimSpace(p.Code)
//	p.Name = strings.TrimSpace(p.Name)
//
//}
