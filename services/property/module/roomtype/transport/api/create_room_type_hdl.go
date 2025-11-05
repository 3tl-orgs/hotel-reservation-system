package roomtypeapi

import (
	"github.com/gin-gonic/gin"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (r *RoomTypeTransport) CreateRoomTypeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data roomtypemodel.RoomTypeCreateDTO

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		if err := data.Validate(); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		if err := r.roomTypeBusiness.CreateRoomTypeBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))
	}
}
