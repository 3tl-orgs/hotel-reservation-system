package propertystore

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func NewPropertyStore(db *gorm.DB) *store {
	return &store{
		db: db,
	}
}
