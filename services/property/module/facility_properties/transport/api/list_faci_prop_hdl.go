package facilitypropertiesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
)

func (fp *FacilityPropertiesTransport) ListFacilityPropHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := fp.transport.ListFacilityPropBiz(c)
		if err != nil {

			core.WriteErrorResponse(c, err)
		}
		
		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
