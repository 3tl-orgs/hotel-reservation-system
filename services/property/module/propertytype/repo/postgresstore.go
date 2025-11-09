package propertytyperepo

import "gorm.io/gorm"

type postgresStore struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *postgresStore {
	return &postgresStore{db: db}
}
