package propertyapi

import (
	"github.com/gin-gonic/gin"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *propertyApi) CreatePropertyHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req propertymodel.PropertyCreateDTO

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, core.Error(400, "Invalid input: "+err.Error(), nil))
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(400, core.Error(400, err.Error(), nil))
			return
		}

		if err := api.business.CreatePropertyBiz(c, &req); err != nil {
			c.JSON(500, core.Error(500, err.Error(), nil))
			return
		}

		c.JSON(201, core.Success("Property created successfully", true))
	}
}
