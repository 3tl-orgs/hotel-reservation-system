package wardsrepo

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (w *wardRepo) GetByCode(ctx context.Context, code string) (*wardsmodel.Ward, error) {
	var data wardsmodel.Ward
	if err := w.db.Table(wardsmodel.Ward{}.TableName()).
		Where("code = ?", code).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wardsmodel.ErrWardNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
