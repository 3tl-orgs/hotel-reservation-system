package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (w *wardTransport) UpdateWardHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data wardsmodel.UpdateWardDTO

		//id, err := strconv.Atoi(c.Param("id"))
		id, err := core.FromBase58(c.Param("id"))
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

		if err := w.wardBusiness.UpdateWardBiz(c, int(id.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}

}
