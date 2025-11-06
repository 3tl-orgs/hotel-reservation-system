package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceTransport) GetProvinceByIdHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		id, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}

		data, err := p.provinceBusiness.GetProvinceByIdBiz(c, int(id.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
