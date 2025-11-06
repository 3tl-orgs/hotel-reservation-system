package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/composer"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	countryApiTransport := composer.NewComposerCountryApiTransport(serviceCtx)
	provinceApiTransport := composer.NewComposerProvinceApiTransport(serviceCtx)
	wardApiTransport := composer.NewComposerWardApiTransport(serviceCtx)

	locations := router.Group("/locations")
	{
		// countries
		locations.POST("/countries", countryApiTransport.CreateCountryHdl())
		locations.GET("/countries/:id", countryApiTransport.GetCountryByIdHdl())
		locations.PUT("/countries/:id", countryApiTransport.UpdateCountryHdl())
		locations.GET("/countries", countryApiTransport.ListCountryHdl())
		locations.GET("/countries/code/:code", countryApiTransport.GetCountryByIdHdl())
		locations.DELETE("/countries/:id", countryApiTransport.DeleteCountryHdl())

		// provinces
		locations.GET("/provinces/:id", provinceApiTransport.GetProvinceByIdHdl())
		locations.POST("/provinces", provinceApiTransport.CreateProvinceHdl())
		locations.GET("/provinces/code/:code", provinceApiTransport.GetProvinceByCodeHdl())
		locations.PUT("/provinces/:id", provinceApiTransport.UpdateProvinceHdl())
		locations.DELETE("/provinces/:id", provinceApiTransport.DeleteProvinceHdl())
		locations.GET("/provinces", provinceApiTransport.ListProvinceHdl())

		// wards
		locations.GET("/wards/:id", wardApiTransport.GetWardByIdHdl())
		locations.POST("/wards", wardApiTransport.CreateWardHdl())
		locations.GET("wards/code/:code", wardApiTransport.GetWardByCodeHdl())
		locations.PUT("/wards/:id", wardApiTransport.UpdateWardHdl())
		locations.DELETE("/wards/:id", wardApiTransport.DeleteWardHdl())
		locations.GET("/wards", wardApiTransport.ListWardHdl())
	}
}
