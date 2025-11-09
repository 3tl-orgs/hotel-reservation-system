package propertybiz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

func (b *business) CreatePropertyBiz(ctx context.Context, req *propertymodel.PropertyCreateDTO) error {
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

	// thumbnail, err := HandleImage(ctx, core.FolderPropertyImage, req.ThumbnailFile, b.uploadProvider)

	// if err != nil {
	// 	return core.ErrInternalServerError.WithError(err.Error())
	// }

	// req.Thumbnail = thumbnail

	property := &propertymodel.Property{
		PropertyTypeId: propertyTypeId,
		CountryId:      countryId,
		ProvinceId:     provinceId,
		WardId:         wardId,
		Thumbnail:      nil,
		Address:        req.Address,
		Lat:            req.Lat,
		Lng:            req.Lng,
		Name:           req.Name,
		Star:           nil,
	}

	// images := core.Images{}
	// if req.Images != nil {
	// 	for _, item := range req.ImageFiles {
	// 		image, err := HandleImage(ctx, core.FolderPropertyImage, item, b.uploadProvider)
	// 		if err != nil {
	// 			return core.ErrInternalServerError.WithError(err.Error())
	// 		}
	// 		images = append(images, *image)
	// 	}
	// }

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
	var bestAmenitiesFormat core.JSONTypeInt
	for _, item := range bestAmenities {
		uid, err := core.FromBase58(item)
		log.Print(int(uid.GetLocalID()))
		if err != nil {
			return core.ErrInternalServerError.WithError(err.Error())
		}
		bestAmenitiesFormat = append(bestAmenitiesFormat, int(uid.GetLocalID()))
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

	if err := b.propertyStore.Create(ctx, property); err != nil {
		return core.ErrInternalServerError.WithError(model.CannotCreatePropertyTypeErr.Error()).WithDebug(err.Error())
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

func HandleImage(ctx context.Context,
	folder string,
	fileHeader *multipart.FileHeader,
	up uploadprovider.UploadProvider) (*core.Image, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, core.ErrUploadFile
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, core.ErrUploadFile
	}

	config, format, err := image.DecodeConfig(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, core.ErrUploadFile
	}

	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), format)

	fmt.Printf("Image format: %s, width: %d, height: %d\n", format, config.Width, config.Height)

	contentType := fileHeader.Header.Get("Content-Type")
	dst := filepath.Join(folder, filename)

	img, err := up.SaveFileUploaded(ctx, file, dst, contentType)

	if err != nil {
		return nil, core.ErrUploadFile
	}

	img.Width = config.Width
	img.Height = config.Height

	return img, nil
}
