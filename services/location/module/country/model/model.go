package model

import "github.com/ngleanhvu/go-booking/shared/core"

type Country struct {
	*core.SQLModel `json:",inline"`
	Name           string `json:"name" gorm:"column:name"`
	Code           string `json:"code" gorm:"column:code,unique"`
}

func (Country) TableName() string {
	return "countries"
}
