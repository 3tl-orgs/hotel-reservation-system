package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) GetAmenityByIdsHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.FilterIds

		if err := c.ShouldBindJSON(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		var ids = make([]int, len(filter.Ids))

		for i, id := range filter.Ids {
			uid, err := core.FromBase58(id)
			if err != nil {
				core.WriteErrorResponse(c, core.ErrBadRequest.
					WithError(err.Error()).
					WithDebug(err.Error()))
				return
			}
			ids[i] = int(uid.GetLocalID())
		}

		data, err := api.amenityService.GetAmenityByIdsBiz(c, ids)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(200, core.ResponseData(data))
	}
}
