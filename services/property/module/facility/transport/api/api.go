package facilityapi

import (
	"context"
	"io"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type FacilityBusiness interface {
	//Biz
	CreateFacilityBiz(ctx context.Context, data *model.FacilityCreateDto, file io.Reader, folder string) error
	GetFacilityByIdBiz(ctx context.Context, id int) (*model.Facility, error)
	GetListFacilitiesBiz(ctx context.Context, paging *core.Paging, moreKeys ...string,) ([]model.Facility, error)
	UpdateFacilityBiz(ctx context.Context, id int, data *model.FacilityUpdateDto) error
	DeleteFacilityBiz(ctx context.Context, id int) error 
}

type facilityApi struct {
	facilityBusiness FacilityBusiness
}

func NewFacilityApi(facilityBusiness FacilityBusiness) *facilityApi {
	return &facilityApi{facilityBusiness: facilityBusiness}
}
