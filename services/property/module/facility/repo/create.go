package facilityrepo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/pkg/errors"
)

func (s *facilityRepo) Create(ctx context.Context, data *model.FacilityCreateDto) error {

	if err := s.db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
