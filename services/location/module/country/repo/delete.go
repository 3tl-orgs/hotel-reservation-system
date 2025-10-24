package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(model.Country{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": false}).
		Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
