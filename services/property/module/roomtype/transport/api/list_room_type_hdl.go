package roomtypeapi

import (
	"github.com/gin-gonic/gin"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (r *RoomTypeTransport) ListRoomTypeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging core.Paging
		if err := c.ShouldBindQuery(&paging); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		paging.Fulfill()

		var filter roomtypemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		data, err := r.roomTypeBusiness.ListRoomTypeBiz(c, &paging, &filter)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError(roomtypemodel.ErrCannotGetRoom.Error()))
		}

		c.JSON(http.StatusOK, core.SuccessResponse(data, paging, filter))
	}
}
