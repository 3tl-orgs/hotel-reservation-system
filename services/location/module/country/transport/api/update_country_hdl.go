package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *api) UpdateCountryHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		var data model.CountryUpdateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		if err := data.Validate(); err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		if err := api.business.UpdateCountryBiz(c, int(uid.GetLocalID()), &data); err != nil {
			c.JSON(500, core.Error(500, err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, core.Success("", data))
	}
}
