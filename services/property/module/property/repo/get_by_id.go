package propertyrepo

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (p *propertyRepo) GetById(ctx context.Context, id int) (*propertymodel.Property, error) {
	var data propertymodel.Property
	if err := p.db.Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(propertymodel.ErrCannotGetProperty)
		}
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
