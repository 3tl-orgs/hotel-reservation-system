package provincebusiness

import (
	"context"
	"fmt"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (p *provinceBusiness) UpdateProvinceBiz(ctx context.Context, id int, data *provincemodel.UpdateProvinceDTO) error {
	existingData, err := p.GetProvinceByIdBiz(ctx, id)

	if err != nil {
		return provincemodel.ErrProvinceNotFound
	}
	if existingData.Id != id {
		return provincemodel.ErrProvinceNotFound
	}
	fmt.Println("Id cua Province: ", id)
	// existingData != nil ==> Có dữ liệu
	// err == nil

	if err := p.provinceRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotUpdateProvince.Error())
	}

	timeUpdated := time.Now()
	data.UpdatedAt = &timeUpdated
	data.Mask()

	return nil
}
