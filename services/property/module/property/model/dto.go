package propertymodel

import (
	"mime/multipart"
	"strings"

	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

type PropertyCreateDTO struct {
	Name           string                  `form:"name" binding:"required"`
	Address        string                  `form:"address" binding:"required"`
	CountryId      string                  `form:"country_id" binding:"required"`
	ProvinceId     string                  `form:"province_id" binding:"required"`
	WardId         string                  `form:"ward_id" binding:"required"`
	ThumbnailFile  *multipart.FileHeader   `form:"thumbnail"`
	Thumbnail      *core.Image             `gorm:"column:thumbnail"`
	Lat            *float64                `form:"lat"`
	Lng            *float64                `form:"lng"`
	PropertyTypeId string                  `form:"property_type_id"`
	Description    *string                 `form:"description"`
	ImageFiles     []*multipart.FileHeader `form:"image_files"` // ✅ đổi sang slice thường
	Images         *core.Images            `gorm:"column:images"`
	CheckInTime    *string                 `form:"check_in_time"`
	CheckoutTime   *string                 `form:"check_out_time"`
	Hotline        *string                 `form:"hotline"`
	HostId         *int                    `form:"host_id"`
	WebsiteUrl     *string                 `form:"website_url"`
	Email          *string                 `form:"email"`
	BestAmenities  *string                 `form:"best_amenities"`
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
