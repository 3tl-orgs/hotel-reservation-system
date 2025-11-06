package facilityapi

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) CreateFacilityHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.FacilityCreateDto
		
		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		fileHeader, _ := c.FormFile("icon")
		var file io.Reader
		if fileHeader != nil {
			f, _ := fileHeader.Open()
			file = f
			defer f.Close()
		}

		folder := "facility-icons/"

		if err := api.facilityBusiness.CreateFacilityBiz(c.Request.Context(), &data, file, folder); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusCreated, core.ResponseData(data))
	}

}
