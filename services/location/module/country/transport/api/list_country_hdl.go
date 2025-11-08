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
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		pagingData.Fulfill()

		var filter model.Filter

		if err := c.ShouldBindJSON(&filter); err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		log.Println("filter:", filter)

		data, err := api.business.ListCountryBiz(c, &filter, &pagingData)

		if err != nil {
			c.JSON(500, core.Error(500, err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, core.Success("", data))
	}
}
