package postgres

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/pkg/errors"
)

func (s *postgresRepo) Update(ctx context.Context, id int, data *model.AmenityUpdateDto) error {
	updateData := map[string]interface{}{}

	if data.Name != nil {
		updateData["name"] = *data.Name
	}

	if data.Icon != nil {
		updateData["icon"] = *data.Icon
	}

	if data.UpdatedAt != nil {
		updateData["updated_at"] = *data.UpdatedAt
	}

	if err := s.db.WithContext(ctx).
		Table(model.Amenity{}.TableName()).
		Where("id = ? AND status = ?", id, true).
		Updates(updateData).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
