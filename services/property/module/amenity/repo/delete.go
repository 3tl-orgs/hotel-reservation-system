package amenityrepo

import (
	"context"

	"time"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) Delete(ctx context.Context, id int) error {

	now := time.Now()

	if err := s.db.WithContext(ctx).
		Table(amenitymodel.Amenity{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": now,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
