package propertytypebiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) CreatePropertyTypeBiz(ctx context.Context,
	data *propertytypemodel.PropertyTypeCreateDto) error {

	isExistByName := b.propertyTypeRepo.ExistByName(ctx, data.Name)

	if isExistByName {
		return core.ErrInternalServerError.WithError(propertytypemodel.PropertyTypeNameIsDuplicateErr.Error())
	}

	if err := b.propertyTypeRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(propertytypemodel.CannotCreatePropertyTypeErr.Error())
	}

	return nil
}
