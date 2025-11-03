package facilitypropertiesrepo

import "gorm.io/gorm"

type facilityPropertiesRepo struct {
	db *gorm.DB
}

func NewFacilityPropertiesRepo(db *gorm.DB) *facilityPropertiesRepo {
	return &facilityPropertiesRepo{
		db: db,
	}
}
