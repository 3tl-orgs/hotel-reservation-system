package amenityrepo

import (
	"context"
	"time"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) DeleteMany(ctx context.Context, ids []int) error {
	now := time.Now()

	if err := s.db.WithContext(ctx).
		Table(amenitymodel.Amenity{}.TableName()).
		Where("id IN ? AND status = ?", ids, true).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": now,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
