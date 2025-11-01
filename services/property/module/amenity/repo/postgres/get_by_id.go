package postgres

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *postgresRepo) GetById(ctx context.Context, id int) (*model.Amenity, error) {

	var amenity model.Amenity

	if err := s.db.WithContext(ctx).Table(model.Amenity{}.TableName()).
		Where("id = ? AND status = ?", id, true).First(&amenity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &amenity, nil
}
