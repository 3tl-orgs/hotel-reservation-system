package amenitybiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/dto"
)

func (b *business) TestGrpcBiz(ctx context.Context, id int32) (*dto.CountryResponse, error) {
	data, err := b.countryRepo.GetCountryById(ctx, id)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error()).WithDebug(err.Error())
	}
	return data, nil
}
