package facilitypropertiesapi

import (
	"github.com/gin-gonic/gin"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (fp *FacilityPropertiesTransport) CreateFacilityPropHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data facilitypropertiesmodel.FacilityPropertiesCreateDTO

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		// Marshal
		//jSon, _ := json.Marshal(body.AmenityIds)
		//
		//data := facilitypropertiesmodel.FacilityPropertiesCreateDTO{
		//	PropertyId: body.PropertyId,
		//	FacilityId: body.FacilityId,
		//	AmenityIds: datatypes.JSON(jSon),
		//}

		if err := fp.transport.CreateFacilityPropBiz(c, &data); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		data.Mask()
		c.JSON(http.StatusCreated, core.ResponseData(data.FakeId))

	}
}
