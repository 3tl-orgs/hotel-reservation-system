package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *propertyTypeApi) GetPropertyTypeByIdHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(400, core.Error(400, "", core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error())))
		}

		data, err := api.propertyService.GetPropertyTypeByIdBiz(c, int(id.GetLocalID()))

		if err != nil {
			c.JSON(500, core.Error(500, "", err))
		}

		c.JSON(200, core.Success("", data))
	}
}
