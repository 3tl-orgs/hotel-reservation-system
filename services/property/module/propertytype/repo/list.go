package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresStore) List(
	ctx context.Context,
	filter *model.Filter,
	paging core.Paging,
	moreInfo ...string,
) ([]model.PropertyType, error) {

	db := s.db.WithContext(ctx).Model(&model.PropertyType{})

	// Preload quan hệ (nếu có)
	for _, m := range moreInfo {
		db = db.Preload(m)
	}

	// Hàm modify cho Paginate
	modify := func(q *gorm.DB) *gorm.DB {
		if filter != nil && filter.Keyword != "" {
			q = q.Where("name ILIKE ?", "%"+filter.Keyword+"%")
		}
		return q
	}

	// Các trường cho phép sắp xếp để tránh SQL injection
	allowedSortFields := []string{"id", "name", "created_at", "updated_at"}

	// Gọi hàm phân trang dùng chung
	data, err := core.Paginate[model.PropertyType](db, &paging, allowedSortFields, modify)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
