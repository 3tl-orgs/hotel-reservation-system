package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *propertyTypeApi) ListPropertyTypeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter
		if err := c.ShouldBindJSON(&filter); err != nil {
			c.JSON(http.StatusBadRequest, core.Error(400, "", core.ErrBadRequest.
				WithError(err.Error())))
		}

		var paging core.Paging
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, core.Error(400, "", core.ErrBadRequest.
				WithError(err.Error())))
		}

		paging.Fulfill()

		data, err := api.propertyService.ListPropertyTypeBiz(c, &filter, paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, core.Error(500, "", err))
		}

		for i := range data {
			data[i].Mask()
		}

		c.JSON(http.StatusOK, core.Success("", data))
	}
}
