package propertyapi

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (p *propertyTransport) DeletePropertyHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()))
			return
		}

		if err := p.propertyBusiness.DeletePropertyBiz(c, int(id.GetLocalID())); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
