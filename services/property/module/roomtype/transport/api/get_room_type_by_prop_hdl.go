package roomtypeapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
	"strconv"
)

func (r *RoomTypeTransport) GetRoomTypeByPropHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		fmt.Println(">>> Param id: " + idStr)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}
		fmt.Println("After parsing Id")

		data, err := r.roomTypeBusiness.GetRoomTypeByProp(c, id)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))

	}
}
