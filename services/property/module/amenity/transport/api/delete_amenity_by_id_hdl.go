package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) DeleteAmenityByIdHdl() gin.HandlerFunc {
	return func (c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		if err := api.amenityService.DeleteAmenityByIdBiz(c, int(uid.GetLocalID())); err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(204, core.ResponseData(true))
	}
}