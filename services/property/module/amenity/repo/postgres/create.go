package postgres

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) Create(ctx context.Context, data *model.AmenityCreateDto) error {
	if err := s.db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
