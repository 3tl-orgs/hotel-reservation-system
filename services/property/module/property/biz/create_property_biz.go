package propertybiz

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *propertyBusiness) CreatePropertyBiz(ctx context.Context, data *propertymodel.PropertyCreateDTO) error {
	if err := p.propertyRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(propertymodel.ErrCannotCreateProperty.Error()).
			WithDebug(err.Error())
	}

	return nil
}
