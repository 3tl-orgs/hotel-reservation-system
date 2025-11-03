package wardsmodel

import (
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type Ward struct {
	core.SQLModel
	Name       string                  `json:"name" gorm:"column:name"`
	Code       string                  `json:"code" gorm:"column:code;unique"`
	ProvinceId int                     `json:"-" gorm:"column:province_id"`
	Province   *provincemodel.Province `json:"province,omitempty" gorm:"preload:false"`
}

func (Ward) TableName() string {
	return "wards"
}

func (ward *Ward) Mask() {
	ward.SQLModel.Mask(core.MaskTypeWard)

	if u := ward.Province; u != nil {
		u.Mask()
	}
}
