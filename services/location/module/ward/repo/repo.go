package wardsrepo

import "gorm.io/gorm"

type wardRepo struct {
	db *gorm.DB
}

func NewWardRepo(db *gorm.DB) *wardRepo {
	return &wardRepo{
		db: db,
	}
}
