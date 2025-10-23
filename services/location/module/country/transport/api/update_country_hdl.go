package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (api *api) UpdateCountryHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		var data model.CountryUpdateDto

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		if err := data.Validate(); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := api.business.UpdateCountryBiz(c, int(uid.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
