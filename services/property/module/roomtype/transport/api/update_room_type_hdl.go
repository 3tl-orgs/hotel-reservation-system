package roomtypeapi

import (
	"github.com/gin-gonic/gin"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (r *RoomTypeTransport) UpdateRoomTypeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data roomtypemodel.RoomTypeUpdateDTO

		id, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrNotFound.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		if err := r.roomTypeBusiness.UpdateRoomTypeBiz(c, int(id.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
