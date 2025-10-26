package provincetransport

import (
	"github.com/gin-gonic/gin"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
	"strconv"
)

func (p *provinceTransport) UpdateProvinceHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data provincemodel.UpdateProvinceDTO

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrNotFound.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}
		// Validate

		if err := p.provinceBusiness.UpdateProvinceBiz(c, id, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
