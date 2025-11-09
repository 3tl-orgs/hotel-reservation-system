package propertystore

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *store) Create(ctx context.Context,
	data *propertymodel.Property,
	data1 *propertydetailmodel.PropertyDetail) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Create(data).Error; err != nil {
			return errors.WithStack(err)
		}
		data1.PropertyId = data.Id
		if err := s.db.Create(data1).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
