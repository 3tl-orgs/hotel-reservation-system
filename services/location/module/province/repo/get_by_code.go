package provincerepo

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (p *provinceRepo) GetByCode(ctx context.Context, code string) (*provincemodel.Province, error) {
	var data provincemodel.Province

	if err := p.db.Table(provincemodel.Province{}.TableName()).
		Where("code = ?", code).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
