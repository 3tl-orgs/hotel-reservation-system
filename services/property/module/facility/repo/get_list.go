package facilityrepo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (s *facilityRepo) List(ctx context.Context,
	paging *core.Paging,
	moreKeys ...string,
) ([]model.Facility, error) {
	var result []model.Facility

	db := s.db.Table(model.Facility{}.TableName())

	for _, key := range moreKeys {
		db.Preload(key)
	}

	var total int64

	if err := db.Count(&total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	paging.Total = int(total)

	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
