package facilityapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *facilityApi) CreateFacilityHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.FacilityCreateDto
		
		// Lấy từng field từ form
		data.Name = c.PostForm("name")
		
		if data.Name == "" {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError("name is required"))
			return
		}

		fileHeader, err := c.FormFile("icon")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError("icon file is required").
				WithDebug(err.Error()))
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError("cannot open file").
				WithDebug(err.Error()))
			return
		}
		defer file.Close()
		
		contentType := fileHeader.Header.Get("Content-Type")

		folder := "facility-icons/"

		if err := api.facilityBusiness.CreateFacilityBiz(c.Request.Context(), &data, file, folder, contentType); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusCreated, core.ResponseData(data))
	}
}
