package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresRepo) Create(ctx context.Context, data *model.CountryCreateDto) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Table(data.TableName()).Create(data).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
