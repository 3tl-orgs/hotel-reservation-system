package facilityrepo

import (
	"context"
	"time"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/pkg/errors"
)

func (s *facilityRepo) Delete(ctx context.Context, id int) error {

	now := time.Now()

	if err := s.db.WithContext(ctx).
		Table(model.Facility{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": now,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}