package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *amenityApi) DeleteAmenityByIdsHdl() gin.HandlerFunc {
	return func (c *gin.Context)  {
		var filter model.FilterIds

		if err := c.ShouldBindJSON(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError("Invalid request body").
				WithDebug(err.Error()))
			return
		}

		var ids = make([]int, len(filter.Ids)) //Tạo slice gồm [len]
		for i, id := range filter.Ids {
			uid, err := core.FromBase58(id);  
			if err != nil {
				core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()). 
				WithDebug(err.Error()))
				return 
			}
			ids[i] = int(uid.GetLocalID()) //gán vào ids
		}

		if err := api.amenityService.DeleteAmenityByIdsBiz(c, ids); err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(204, core.ResponseData(true))
	}
}