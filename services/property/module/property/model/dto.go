package propertymodel

import (
	"mime/multipart"
	"strings"

	"github.com/pkg/errors"
)

type PropertyCreateDTO struct {
	Name           string                  `form:"name" gorm:"column:name"`
	Address        string                  `form:"address" gorm:"column:address"`
	CountryId      string                  `form:"country_id" gorm:"column:country_id"`
	ProvinceId     string                  `form:"province_id" gorm:"column:province_id"`
	WardId         string                  `form:"ward_id" gorm:"column:ward_id"`
	Thumbnail      *multipart.FileHeader   `form:"thumbnail" gorm:"column:thumbnail"`
	Lat            *float64                `form:"lat" gorm:"column:lat"`
	Lng            *float64                `form:"lng" gorm:"column:lng"`
	PropertyTypeId string                  `form:"property_type_id" gorm:"column:property_type_id"`
	Description    *string                 `form:"description" gorm:"column:description"`
	Image_files    *[]multipart.FileHeader `form:"image_files" gorm:"column:image_files"`
	Images         *[]string               `form:"images" gorm:"column:images"`
	CheckInTime    *string                 `form:"check_in_time" gorm:"column:check_in_time"`
	CheckoutTime   *string                 `form:"checkout_time" gorm:"column:checkout_time"`
	Hotline        *string                 `form:"hotline" gorm:"column:hotline"`
	HostId         *int                    `form:"host_id" gorm:"column:host_id"`
	WebsiteUrl     *string                 `form:"website_url" gorm:"column:website_url"`
	Email          *string                 `form:"email" gorm:"column:email"`
	BestAmenities  *string                 `form:"best_amenities" gorm:"column:best_amenities"`
}

func (PropertyCreateDTO) TableName() string {
	return Property{}.TableName()
}

func (p *PropertyCreateDTO) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	p.Address = strings.TrimSpace(p.Address)
	p.PropertyTypeId = strings.TrimSpace(p.PropertyTypeId)
	p.CountryId = strings.TrimSpace(p.CountryId)
	p.ProvinceId = strings.TrimSpace(p.ProvinceId)
	p.WardId = strings.TrimSpace(p.WardId)

	if err := checkName(p.Name); err != nil {
		return errors.WithStack(err)
	}

	if err := checkPropertyTypeId(p.PropertyTypeId); err != nil {
		return errors.WithStack(err)
	}

	if err := checkAddress(p.Address); err != nil {
		return errors.WithStack(err)
	}

	if err := checkCountryId(p.CountryId); err != nil {
		return errors.WithStack(err)
	}

	if err := checkProvinceId(p.ProvinceId); err != nil {
		return errors.WithStack(err)
	}

	if err := checkWardId(p.WardId); err != nil {
		return errors.WithStack(err)
	}

	return nil

}
