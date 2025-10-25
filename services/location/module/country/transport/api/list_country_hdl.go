package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *api) ListCountryHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData core.Paging

		if err := c.ShouldBindQuery(&pagingData); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		pagingData.Fulfill()

		var filter model.Filter

		if err := c.ShouldBindJSON(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		log.Println("filter:", filter)

		data, err := api.business.ListCountryBiz(c, &filter, &pagingData)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
