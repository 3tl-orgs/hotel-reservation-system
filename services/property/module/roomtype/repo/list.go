package roomtyperepo

import (
	"context"
	"fmt"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (r *roomTypeRepo) List(ctx context.Context,
	paging *core.Paging,
	filter *roomtypemodel.Filter,
	moreKeys ...string,
) ([]roomtypemodel.RoomType, error) {
	var data []roomtypemodel.RoomType

	db := r.db.Table(roomtypemodel.RoomType{}.TableName()).Where("status = ?", true)

	if filter != nil && filter.Keyword != "" {
		keyword := "%" + filter.Keyword + "%"
		db = db.Where("name LIKE ?", keyword)
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
	orderEx := fmt.Sprintf("%s %s", paging.SortBy, paging.SortDir)

	if err := db.Offset(offset).
		Limit(paging.Limit).
		Order(orderEx).
		Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
