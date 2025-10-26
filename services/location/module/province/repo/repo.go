package provincerepo

import "gorm.io/gorm"

type provinceRepo struct {
	db *gorm.DB
}

func NewProvinceRepo(db *gorm.DB) *provinceRepo {
	return &provinceRepo{
		db: db,
	}
}
