package propertybiz

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (p *propertyBusiness) GetPropertyByIdBiz(ctx context.Context, id int) (*propertymodel.Property, error) {
	data, err := p.propertyRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(propertymodel.ErrCannotGetProperty.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(propertymodel.ErrCannotGetProperty.Error())
	}

	data.Mask()
	return data, nil
}
