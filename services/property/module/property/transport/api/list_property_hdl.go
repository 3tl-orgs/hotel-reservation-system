package propertyapi

import (
	"github.com/gin-gonic/gin"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (p *propertyTransport) ListPropertyHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData core.Paging
		if err := c.ShouldBindQuery(&pagingData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(propertymodel.ErrCannotGetProperty.Error()).
				WithDebug(err.Error()),
			)
			return
		}
		pagingData.Fulfill()

		var filter propertymodel.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(propertymodel.ErrCannotGetProperty.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		data, err := p.propertyBusiness.ListPropertyBiz(c, &pagingData, &filter)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError(propertymodel.ErrCannotGetProperty.Error()))
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(data, pagingData, filter))
	}
}
