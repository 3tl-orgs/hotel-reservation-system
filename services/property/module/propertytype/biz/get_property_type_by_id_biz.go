package propertytypebiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) GetPropertyTypeByIdBiz(ctx context.Context, id int) (*propertytypemodel.PropertyType, error) {
	data, err := b.propertyTypeRepo.GetById(ctx, id)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(propertytypemodel.CannotGetPropertyTypeErr.Error())
	}
	return data, nil
}
