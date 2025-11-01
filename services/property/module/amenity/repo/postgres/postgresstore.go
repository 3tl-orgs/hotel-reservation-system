package postgres

import "gorm.io/gorm"

type postgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *postgresRepo {
	return &postgresRepo{db: db}
}
