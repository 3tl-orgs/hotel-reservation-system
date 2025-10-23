package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) Update(ctx context.Context, id int, data *model.CountryUpdateDto) error {
	if err := s.db.Table(model.Country{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
