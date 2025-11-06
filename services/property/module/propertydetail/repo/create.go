package propertydetailstore

import (
	"context"

	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/pkg/errors"
)

func (s *store) Update(ctx context.Context, data *propertydetailmodel.PropertyDetail) error {
	if err := s.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
