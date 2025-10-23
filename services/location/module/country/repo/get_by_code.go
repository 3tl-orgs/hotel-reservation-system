package repo

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresRepo) GetByCode(ctx context.Context, code string) (*model.Country, error) {
	var country model.Country
	if err := s.db.Table(model.Country{}.TableName()).
		Where("code = ?", code).
		First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &country, nil
}
