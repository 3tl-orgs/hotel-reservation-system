package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (w *wardTransport) GetWardByCodeHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		if strings.TrimSpace(code) == "" {
			core.WriteErrorResponse(c, wardsmodel.ErrWardCodeIsEmpty)
			return
		}

		data, err := w.wardBusiness.GetWardByCodeBiz(c, code)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
