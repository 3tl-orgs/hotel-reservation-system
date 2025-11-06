package propertydetailstore

import (
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func NewPropertyDetailStore(db *gorm.DB) *store {
	return &store{db: db}
}
