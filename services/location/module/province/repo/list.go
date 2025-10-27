package provincerepo

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"strings"
)

func (p *provinceRepo) ListData(
	ctx context.Context,
	filter *provincemodel.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]provincemodel.Province, error) {
	var result []provincemodel.Province

	db := p.db.Table(provincemodel.Province{}.TableName()).Where("status is true")

	if f := filter; f != nil {
		keyword := strings.TrimSpace(f.Keyword)
		if keyword == "" {
			db = db.Where("name LIKE ?", "%"+keyword+"%")
		}
	}

	for _, key := range moreKeys {
		db.Preload(key)
	}
	var count int64

	if err := db.Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	paging.Total = int(count)
	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
