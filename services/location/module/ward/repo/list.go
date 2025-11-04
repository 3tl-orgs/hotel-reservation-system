package wardsrepo

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"strings"
)

func (w *wardRepo) ListData(
	ctx context.Context,
	filter *wardsmodel.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]wardsmodel.Ward, error) {
	var result []wardsmodel.Ward
	db := w.db.Table(wardsmodel.Ward{}.TableName()).Where("status = true")

	if f := filter; f != nil {
		keyword := strings.TrimSpace(f.Keyword)
		if keyword == "" {
			db = db.Where("name LIKE ?", "%"+keyword+"%")
		}
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	paging.Total = int(count)

	for _, key := range moreKeys {
		db.Preload(key)
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := db.Offset(offset).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
