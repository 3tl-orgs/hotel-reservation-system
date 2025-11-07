package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresStore) GetById(ctx context.Context, id int) (*model.PropertyType, error) {
	var data model.PropertyType
	if err := s.db.Table(model.PropertyType{}.TableName()).
		Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
