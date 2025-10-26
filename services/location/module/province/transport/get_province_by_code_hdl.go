package provincetransport

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/shared/core"
	"net/http"
	"strings"
)

func (p *provinceTransport) GetProvinceByCodeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		if strings.TrimSpace(code) == "" {
			core.WriteErrorResponse(c, core.ErrBadRequest)
		}

		data, err := p.provinceBusiness.GetProvinceByCodeBiz(c, code)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
