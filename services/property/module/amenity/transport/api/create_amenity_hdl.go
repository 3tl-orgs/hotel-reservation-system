package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) CreateAmenityHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.AmenityCreateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		if err := api.amenityService.CreateAmenityBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))
	}
}
