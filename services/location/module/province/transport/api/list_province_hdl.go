package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceTransport) ListProvinceHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData core.Paging
		if err := c.ShouldBindQuery(&pagingData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}
		pagingData.Fulfill()

		var filter provincemodel.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}
		data, err := p.provinceBusiness.ListProvinceBiz(c, &filter, &pagingData)
		if err != nil {
			core.WriteErrorResponse(c, err)
		}

		c.JSON(http.StatusOK, core.SuccessResponse(data, pagingData, filter))
	}
}
