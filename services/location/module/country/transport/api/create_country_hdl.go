package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *api) CreateCountryHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.CountryCreateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		if err := data.Validate(); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		if err := api.business.CreateCountryBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(201, core.ResponseData(data.FakeId))
	}
}
