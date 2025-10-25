package repo

import (
	"context"
	"strings"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (s *postgresRepo) List(ctx context.Context,
	filter *model.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]model.Country, error) {
	var result []model.Country

	db := s.db.Table(model.Country{}.TableName())

	for _, key := range moreKeys {
		db.Preload(key)
	}

	if f := filter; f != nil {
		keyword := strings.TrimSpace(f.Keyword)
		if keyword != "" {
			db = db.Where("lower(name) LIKE ?", "%"+keyword+"%")
		}
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
