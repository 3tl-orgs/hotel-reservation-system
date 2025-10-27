package provincebusiness

import (
	"context"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceBusiness) ListProvinceBiz(ctx context.Context,
	filter *provincemodel.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]provincemodel.Province, error) {
	result, err := p.provinceRepo.ListData(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetCountry.Error()).
			WithDebug(err.Error())
	}

	for i := range result {
		result[i].Mask()
	}
	return result, nil
}
