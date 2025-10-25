package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) GetByIds(ctx context.Context,
	ids []int) ([]model.Amenity, error) {
	var result []model.Amenity
	if err := s.db.Table(model.Amenity{}.TableName()).
		Where("id IN ?", ids).
		Where("status = ?", true).
		Find(result).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}
