package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *api) GetCountryByCodeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		data, err := api.business.GetCountryByCodeBiz(c, code)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(200, core.ResponseData(data))
	}
}
