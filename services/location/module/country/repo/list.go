package repo

import (
	"context"
	"fmt"
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

	db = db.Where("status = ?", true)

	if f := filter; f != nil {
		keyword := strings.TrimSpace(strings.ToLower(f.Keyword))
		if keyword != "" {
			db = db.Where("lower(name) LIKE ?", "%"+keyword+"%")
		}
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	paging.Total = int(total)

	if err := db.Order(fmt.Sprintf("%s %s", paging.SortBy, paging.SortDir)).
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
