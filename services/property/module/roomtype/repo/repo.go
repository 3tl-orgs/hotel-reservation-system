package roomtyperepo

import "gorm.io/gorm"

type roomTypeRepo struct {
	db *gorm.DB
}

func NewRoomTypeRepo(db *gorm.DB) *roomTypeRepo {
	return &roomTypeRepo{
		db: db,
	}
}
