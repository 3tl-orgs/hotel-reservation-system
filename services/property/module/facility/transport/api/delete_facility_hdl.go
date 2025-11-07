package facilityapi

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) DeleteFacilityHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid, err := core.FromBase58(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
		}
		
		if err := api.facilityBusiness.DeleteFacilityBiz(c, id); err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(204, core.ResponseData(true))
	}

}