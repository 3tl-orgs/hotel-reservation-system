package propertystore

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
)

func (s *store) Create(ctx context.Context, data *propertymodel.Property) error {
	if err := s.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
