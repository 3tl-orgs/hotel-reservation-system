package core

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Paginate là hàm generic dùng chung cho mọi model
func Paginate[T any](db *gorm.DB, paging *Paging, allowedSortFields []string, modify func(*gorm.DB) *gorm.DB) ([]T, error) {
	var data []T
	var count int64

	// Chuẩn hóa dữ liệu phân trang
	paging.Fulfill()

	// Kiểm tra field được phép sort (nếu có)
	if len(allowedSortFields) > 0 {
		paging.ValidateSortFields(allowedSortFields)
	}

	// Áp dụng các filter tùy chỉnh nếu có
	if modify != nil {
		db = modify(db)
	}

	// Đếm tổng số bản ghi
	if err := db.Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	paging.Total = int(count)

	// Tính offset
	offset := (paging.Page - 1) * paging.Limit

	// Thêm sắp xếp và phân trang
	db = db.Order(fmt.Sprintf("%s %s", paging.SortBy, paging.SortDir)).
		Offset(offset).
		Limit(paging.Limit)

	// Truy vấn dữ liệu
	if err := db.Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
