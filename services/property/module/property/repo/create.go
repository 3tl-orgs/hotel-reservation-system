package propertystore

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *store) Create(ctx context.Context, data *propertymodel.Property) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Create(data).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
