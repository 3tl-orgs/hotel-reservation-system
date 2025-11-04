package propertyrepo

import "gorm.io/gorm"

type propertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *propertyRepo {
	return &propertyRepo{
		db: db,
	}
}
