package wardstransport

import (
	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
	"strconv"
)

func (w *wardTransport) GetWardByIdHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(wardsmodel.ErrCannotGetWard.Error()),
			)
			return
		}

		data, err := w.wardBusiness.GetWardByIdBiz(c, id)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
