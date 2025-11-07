package facilityapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) GetListFacilitiesHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData core.Paging

		if err := c.ShouldBindQuery(&pagingData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		pagingData.Fulfill()

		data, err := api.facilityBusiness.GetListFacilitiesBiz(c, &pagingData)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
