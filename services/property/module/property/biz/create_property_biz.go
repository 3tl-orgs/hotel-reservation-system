package propertybiz

import (
	"context"
	"fmt"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) CreatePropertyBiz(ctx context.Context, req *propertymodel.PropertyCreateDTO) error {
	// Chuyển UID thành ID int
	propertyTypeId, err := transferData(req.PropertyTypeId)
	if err != nil {
		return err
	}
	countryId, err := transferData(req.CountryId)
	if err != nil {
		return err
	}
	provinceId, err := transferData(req.ProvinceId)
	if err != nil {
		return err
	}
	wardId, err := transferData(req.WardId)
	if err != nil {
		return err
	}

	property := &propertymodel.Property{
		PropertyTypeId: propertyTypeId,
		CountryId:      countryId,
		ProvinceId:     provinceId,
		WardId:         wardId,
		Thumbnail:      req.Thumbnail,
		Address:        req.Address,
		Lat:            req.Lat,
		Lng:            req.Lng,
		Name:           req.Name,
		Star:           nil,
	}

	images := core.JSONType[core.Image]{}
	if req.Images != nil {
		images = *req.Images
	}

	bestAmenities := core.JSONType[int]{}
	if req.BestAmenities != nil && *req.BestAmenities != "" {
		arr, err := core.FromJSON[[]int](*req.BestAmenities)
		if err != nil {
			return fmt.Errorf("invalid bestAmenities json: %w", err)
		}
		bestAmenities = arr
	}

	propertyDetail := &propertydetailmodel.PropertyDetail{
		PropertyId:    property.Id,
		CheckInTime:   req.CheckInTime,
		CheckOutTime:  req.CheckoutTime,
		Email:         req.Email,
		Description:   req.Description,
		Hotline:       req.Hotline,
		HostId:        req.HostId,
		WebsiteUrl:    req.WebsiteUrl,
		Images:        &images,
		BestAmenities: &bestAmenities,
	}

	if err := b.propertyStore.Create(ctx, property); err != nil {
		return core.ErrInternalServerError.WithError(propertymodel.ErrCannotCreateProperty.Error())
	}

	if err := b.propertyDetailStore.Create(ctx, propertyDetail); err != nil {
		return core.ErrInternalServerError.WithError(propertymodel.ErrCannotCreateProperty.Error())
	}

	return nil
}

func transferData(uid string) (int, error) {
	id, err := core.FromBase58(uid)
	if err != nil {
		return 0, core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}
	return int(id.GetLocalID()), nil
}
