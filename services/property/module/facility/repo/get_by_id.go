package facilityrepo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *facilityRepo) GetById(ctx context.Context, id int) (*model.Facility, error) {
	
	var facility model.Facility

	if err := s.db.WithContext(ctx).Table(model.Facility{}.TableName()).
		Where("id = ? AND status = ?", id, true).First(&facility).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &facility, nil
}