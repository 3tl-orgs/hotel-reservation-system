package facilityapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) GetFacilityByIdHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest. 
				WithError(err.Error()). 
				WithDebug(err.Error()))
			return 
		}

		data, err := api.facilityBusiness.GetFacilityByIdBiz(c,id)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(200, core.SuccessResponse(data, nil, nil))
	}
}
