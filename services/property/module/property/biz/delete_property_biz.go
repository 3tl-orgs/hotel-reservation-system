package propertybiz

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *propertyBusiness) DeletePropertyBiz(ctx context.Context, id int) error {
	if err := p.propertyRepo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(propertymodel.ErrCannotDeleteProperty.Error())
	}
	return nil
}
