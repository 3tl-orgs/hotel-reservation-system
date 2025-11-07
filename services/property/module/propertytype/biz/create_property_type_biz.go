package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) CreatePropertyTypeBiz(ctx context.Context,
	data *model.PropertyTypeCreateDto) error {

	isExistByName := b.propertyTypeRepo.ExistByName(ctx, data.Name)

	if isExistByName {
		return core.ErrInternalServerError.WithError(model.PropertyTypeNameIsDuplicateErr.Error())
	}

	if err := b.propertyTypeRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(model.CannotCreatePropertyTypeErr.Error())
	}

	return nil
}
