package api

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
)

type AmenityService interface {
	GetAmenityByIdsBiz(ctx context.Context, ids []int) ([]model.Amenity, error)
	CreateAmenityBiz(ctx context.Context, data *model.AmenityCreateDto) error
	DeleteAmenityByIdsBiz(ctx context.Context, ids []int) error 
	GetAmenityByIdBiz(ctx context.Context, id int) (*model.Amenity, error)
	UpdateAmenityBiz(ctx context.Context, id int, data *model.AmenityUpdateDto) error
	DeleteAmenityByIdBiz(ctx context.Context, id int) error
}

type amenityApi struct {
	amenityService AmenityService
}

func NewAmenityApi(amenityService AmenityService) *amenityApi {
	return &amenityApi{amenityService: amenityService}
}
