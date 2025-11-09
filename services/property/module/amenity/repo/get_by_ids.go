package amenityrepo

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) GetByIds(ctx context.Context,
	ids []int) ([]amenitymodel.Amenity, error) {
	var result []amenitymodel.Amenity
	if err := s.db.Table(amenitymodel.Amenity{}.TableName()).
		Where("id IN ?", ids).
		Where("status = ?", true).
		Find(result).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}
