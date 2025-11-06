package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceTransport) CreateProvinceHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data provincemodel.CreateProvinceDTO

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
		}

		if err := p.provinceBusiness.CreateProvinceBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))
	}
}
