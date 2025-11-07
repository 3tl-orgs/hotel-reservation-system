package facilityrepo

import "gorm.io/gorm"

type facilityRepo struct {
	db *gorm.DB
}

func NewFacilityRepo(db *gorm.DB) *facilityRepo {
	return &facilityRepo{db : db}
}