package propertyapi

import (
	"github.com/gin-gonic/gin"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (p *propertyTransport) UpdatePropertyHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		var data propertymodel.PropertyUpdateDTO
		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		if err := p.propertyBusiness.UpdatePropertyBiz(c, int(id.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
