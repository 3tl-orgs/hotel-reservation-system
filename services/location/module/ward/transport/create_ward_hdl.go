package wardstransport

import (
	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (w *wardTransport) CreateWardHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data wardsmodel.CreateWardDTO

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := w.wardBusiness.CreateWardBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))
	}
}
