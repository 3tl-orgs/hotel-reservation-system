package propertytyperepo

import (
	"context"

	propertytypemodel "github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresStore) GetById(ctx context.Context, id int) (*propertytypemodel.PropertyType, error) {
	var data propertytypemodel.PropertyType
	if err := s.db.Table(propertytypemodel.PropertyType{}.TableName()).
		Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
