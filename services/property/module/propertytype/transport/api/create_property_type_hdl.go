package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *propertyTypeApi) CreatePropertyTypeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.PropertyTypeCreateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, core.Error(400, "",
				core.ErrBadRequest.
					WithError(err.Error()).
					WithDebug(err.Error())))
		}

		if err := api.propertyService.CreatePropertyTypeBiz(c, &data); err != nil {
			c.JSON(500, core.Error(500, "", err))
		}

		c.JSON(200, core.Success("", nil))
	}
}
