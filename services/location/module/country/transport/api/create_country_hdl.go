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
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		if err := data.Validate(); err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		if err := api.business.CreateCountryBiz(c, &data); err != nil {
			c.JSON(500, core.Error(500, err.Error(), nil))
			return
		}

		data.Mask()
		c.JSON(201, core.Success("", data.FakeId))
	}
}
