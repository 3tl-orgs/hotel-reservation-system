package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) UpdateAmenityHdl() gin.HandlerFunc {
	return func (c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id")); 
		
		if err != nil { //Xử lý format, validate ID truyền vào
			core.WriteErrorResponse(c, core.ErrBadRequest.
			WithError(err.Error()). 
			WithDebug(err.Error()))

			return;
		}

		var data model.AmenityUpdateDto

		if err := c.ShouldBindJSON(&data); err != nil { //truyền data vào struct
			core.WriteErrorResponse(c, core.ErrBadRequest.
			WithError(err.Error()). 
			WithDebug(err.Error()))
			return;
		}

		if err := data.Validate(); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return 
		}

		// ctx := c.Request.Context()
		if err := api.amenityService.UpdateAmenityBiz(c, int(uid.GetLocalID()), &data); err != nil {
			// int(uid.GetLocalID()) để decode lấy id thật
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}