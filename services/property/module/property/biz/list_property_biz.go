package propertybiz

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *propertyBusiness) ListPropertyBiz(ctx context.Context,
	pagingData *core.Paging,
	filter *propertymodel.Filter,
	moreKeys ...string,
) ([]propertymodel.Property, error) {

	data, err := p.propertyRepo.List(ctx, pagingData, filter, moreKeys...)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(propertymodel.ErrCannotGetProperty.Error())
	}

	for i := range data {
		data[i].Mask()
	}

	return data, nil
}
