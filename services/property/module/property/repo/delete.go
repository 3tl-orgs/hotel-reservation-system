package propertyrepo

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
	"time"
)

func (p *propertyRepo) Delete(ctx context.Context, id int) error {
	if err := p.db.Table(propertymodel.Property{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
