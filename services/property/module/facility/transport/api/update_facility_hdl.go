package facilityapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) UpdateFacilityHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil { //Xử lý format, validate ID truyền vào
			core.WriteErrorResponse(c, core.ErrBadRequest.
			WithError(err.Error()). 
			WithDebug(err.Error()))

			return;
		}

		var data model.FacilityUpdateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
			WithError(err.Error()). 
			WithDebug(err.Error()))
			return;
		}

		if err := data.Validate(); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return 
		}

		if err := api.facilityBusiness.UpdateFacilityBiz(c, id, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return 
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}

}
