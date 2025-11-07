package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) ListPropertyTypeBiz(ctx context.Context, filter *model.Filter,
	paging core.Paging, moreInfo ...string) ([]model.PropertyType, error) {
	data, err := b.propertyTypeRepo.List(ctx, filter, paging, moreInfo...)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.CannotGetPropertyTypeErr.Error())
	}
	return data, nil
}
