package propertyrepo

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
)

func (p *propertyRepo) Create(ctx context.Context, data *propertymodel.PropertyCreateDTO) error {
	if err := p.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	
	return nil
}
