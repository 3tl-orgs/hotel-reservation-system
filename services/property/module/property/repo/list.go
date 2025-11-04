package propertyrepo

import (
	"context"
	"fmt"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (p *propertyRepo) List(ctx context.Context,
	pagingData *core.Paging,
	filter *propertymodel.Filter,
	moreKeys ...string,
) ([]propertymodel.Property, error) {
	var data []propertymodel.Property

	db := p.db.Table(propertymodel.Property{}.TableName()).Where("status = ?", true)

	if filter != nil && filter.Keyword != "" {
		keyword := "%" + filter.Keyword + "%"
		db = db.Where("name LIKE ? OR address LIKE ?", keyword, keyword)
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	pagingData.Total = int(count)

	for _, key := range moreKeys {
		db.Preload(key)
	}

	fmt.Println("Starting offset")
	offset := (pagingData.Page - 1) * pagingData.Limit
	orderExpr := fmt.Sprintf("%s %s", pagingData.SortBy, pagingData.SortDir)

	if err := db.Offset(offset).
		Limit(pagingData.Limit).
		Order(orderExpr).
		Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
