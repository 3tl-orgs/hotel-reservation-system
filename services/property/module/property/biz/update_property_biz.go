package propertybiz

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (p *propertyBusiness) UpdatePropertyBiz(ctx context.Context, id int, data *propertymodel.PropertyUpdateDTO) error {
	
	if err := p.propertyRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(propertymodel.ErrCannotUpdateProperty.Error())
	}

	updatedTime := time.Now()
	data.UpdatedAt = &updatedTime

	return nil
}
