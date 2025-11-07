package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/pkg/errors"
)

func (s *postgresStore) Create(ctx context.Context,
	data *model.PropertyTypeCreateDto) error {
	if err := s.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
