package propertyapi

import (
	"github.com/gin-gonic/gin"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (p *propertyTransport) CreatePropertyHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data propertymodel.PropertyCreateDTO

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := p.propertyBusiness.CreatePropertyBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		data.Mask()
		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))
	}
}
