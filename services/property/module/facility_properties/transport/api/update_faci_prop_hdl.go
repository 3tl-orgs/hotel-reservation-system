package facilitypropertiesapi

import (
	"github.com/gin-gonic/gin"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (fp *FacilityPropertiesTransport) UpdateFacilityPropHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data facilitypropertiesmodel.FacilityPropertiesUpdateDTO
		id, err := core.FromBase58(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}

		if err := fp.transport.UpdateFacilityPropBiz(c, int(id.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()),
			)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
