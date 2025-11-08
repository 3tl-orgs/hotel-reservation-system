package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresRepo) Delete(ctx context.Context, id int) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Table(model.Country{}.TableName()).
			Where("id = ?", id).
			Updates(map[string]interface{}{"status": false}).
			Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
