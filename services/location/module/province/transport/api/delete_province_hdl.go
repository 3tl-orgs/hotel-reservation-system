package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceTransport) DeleteProvinceHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		id, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrNotFound.
				WithError(err.Error()).
				WithDebug(err.Error()),
			)
			return
		}

		if err := p.provinceBusiness.DeleteProvinceBiz(c, int(id.GetLocalID())); err != nil {
			core.WriteErrorResponse(c, core.ErrNotFound.
				WithError(err.Error()).
				WithDebug(err.Error()))
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
