package wardstransport

import (
	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
	"strconv"
)

func (w *wardTransport) UpdateWardHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data wardsmodel.UpdateWardDTO

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		// Validate
		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := w.wardBusiness.UpdateWardBiz(c, int(id), &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}

}
