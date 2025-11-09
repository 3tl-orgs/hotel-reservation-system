package propertybiz

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

func (b *business) CreatePropertyBiz(ctx context.Context, req *propertymodel.PropertyCreateDTO) error {
	propertyTypeId, err := core.TransferData(req.PropertyTypeId)
	if err != nil {
		return err
	}

	_, err = b.propertyTypeStore.GetById(ctx, propertyTypeId)

	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}

	countryId, err := core.TransferData(req.CountryId)
	if err != nil {
		return err
	}

	_, err = b.countryClient.GetCountryById(ctx, int32(countryId))

	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}

	provinceId, err := core.TransferData(req.ProvinceId)
	if err != nil {
		return err
	}

	_, err = b.provinceClient.GetProvinceById(ctx, int32(provinceId))

	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}

	wardId, err := core.TransferData(req.WardId)
	if err != nil {
		return err
	}

	_, err = b.wardClient.GetWardById(ctx, int32(wardId))
	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}

	if req.ThumbnailFile != nil {
		thumbnail, err := uploadprovider.HandleImage(ctx, core.FolderPropertyImage, req.ThumbnailFile, b.uploadProvider)
		req.Thumbnail = thumbnail
		if err != nil {
			return core.ErrInternalServerError.WithError(err.Error())
		}
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

	images := core.Images{}
	if req.ImageFiles != nil && len(req.ImageFiles) > 0 {
		for _, item := range req.ImageFiles {
			handleImage, err := uploadprovider.HandleImage(ctx, core.FolderPropertyImage, item, b.uploadProvider)
			if err != nil {
				return core.ErrInternalServerError.WithError(err.Error())
			}
			images = append(images, *handleImage)
		}
	}

	var bestAmenities []string

	log.Print("req best amenities ", req.BestAmenities)
	if req.BestAmenities != nil && *req.BestAmenities != "" {
		log.Printf("raw BestAmenities: %s", *req.BestAmenities)

		amenitiesStr := *req.BestAmenities

		// Handle double-encoded JSON (unwrap outer quotes)
		var tempStr string
		if err := json.Unmarshal([]byte(amenitiesStr), &tempStr); err == nil {
			// Successfully unwrapped, use the inner string
			amenitiesStr = tempStr
		}

		// Parse the actual JSON array
		if err := json.Unmarshal([]byte(amenitiesStr), &bestAmenities); err != nil {
			return fmt.Errorf("invalid bestAmenities json: %w", err)
		}
	}

	log.Printf("bestAmenities: %+v", bestAmenities)
	// Convert string[] to int[]
	var bestAmenitiesFormat propertydetailmodel.BestAmenities
	for _, item := range bestAmenities {
		uid, err := core.FromBase58(item)
		log.Print(int(uid.GetLocalID()))
		if err != nil {
			return core.ErrInternalServerError.WithError(err.Error())
		}
		bestAmenitiesFormat = append(bestAmenitiesFormat, int(uid.GetLocalID()))
	}

	amenitiesData, err := b.amenityStore.GetByIds(ctx, bestAmenitiesFormat)

	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	if len(amenitiesData) != len(bestAmenitiesFormat) {
		return core.ErrInternalServerError.WithError(core.ErrConflict.Error())
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
		Images:        nil,
		BestAmenities: &bestAmenitiesFormat,
	}

	if err := b.propertyStore.Create(ctx, property, propertyDetail); err != nil {
		return core.ErrInternalServerError.WithError(propertymodel.ErrCannotCreateProperty.Error())
	}

	return nil
}
