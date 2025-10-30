package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) GetAmenityByIdHdl() gin.HandlerFunc {
	return func (c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest. 
				WithError(err.Error()). 
				WithDebug(err.Error()))
			return 
		}

		data, err := api.amenityService.GetAmenityByIdBiz(c, int(uid.GetLocalID()))

		if err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(200, core.SuccessResponse(data, nil, nil))
	}

}