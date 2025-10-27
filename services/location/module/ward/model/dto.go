package wardsmodel

import "github.com/ngleanhvu/go-booking/shared/core"

type CreateWardDTO struct {
	core.SQLModel
	Name       string `json:"name" gorm:"column:name"`
	Code       string `json:"code" gorm:"column:code"`
	ProvinceId int    `json:"province_id" gorm:"column:province_id"`
}

func (CreateWardDTO) TableName() string {
	return "wards"
}

func (ward *CreateWardDTO) Mask() {
	ward.SQLModel.Mask(core.MaskTypeWard)
}

type UpdateWardDTO struct {
	core.SQLModel
	Name       string `json:"name" gorm:"column:name"`
	Code       string `json:"code" gorm:"column:code"`
	ProvinceId int    `json:"province_id" gorm:"column:province_id"`
}

func (UpdateWardDTO) TableName() string {
	return "wards"
}

func (ward *UpdateWardDTO) Mask() {
	ward.SQLModel.Mask(core.MaskTypeWard)
}
