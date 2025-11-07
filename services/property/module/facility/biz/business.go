package facilitybiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

type FacilityRepo interface {
	Create(ctx context.Context, data *model.FacilityCreateDto) error
	GetById(ctx context.Context, id int) (*model.Facility, error)
	Update(ctx context.Context, id int, data *model.FacilityUpdateDto) error
	List(ctx context.Context,
		paging *core.Paging, moreKey ...string) ([]model.Facility, error)
	Delete(ctx context.Context, id int) error
}


type facilityBusiness struct {
	repo  FacilityRepo
	uploader uploadprovider.UploadProvider
}

func NewFacilityBusiness(repo FacilityRepo, uploader uploadprovider.UploadProvider) *facilityBusiness {
	return &facilityBusiness{repo: repo, uploader: uploader}
}
