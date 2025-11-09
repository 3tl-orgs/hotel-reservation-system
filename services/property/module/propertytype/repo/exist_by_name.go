package propertytyperepo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
)

func (s *postgresStore) ExistByName(ctx context.Context, name string) bool {
	var propertyType propertytypemodel.PropertyType
	_ = s.db.WithContext(ctx).Where("name = ?", name).First(&propertyType)
	return propertyType.Id != 0
}
