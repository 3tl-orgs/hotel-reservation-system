package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (w *wardTransport) ListWardHdl() gin.HandlerFunc {
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
		var filter wardsmodel.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		data, err := w.wardBusiness.ListWardBiz(c, &filter, &pagingData)
		if err != nil {
			core.WriteErrorResponse(c, err)
		}

		c.JSON(http.StatusOK, core.SuccessResponse(data, pagingData, filter))
	}
}
