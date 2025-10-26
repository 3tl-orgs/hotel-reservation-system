package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/composer"
	"github.com/ngleanhvu/go-booking/shared/core"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc/middleware"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/gormc"
	"github.com/spf13/cobra"
)

func newServiceCtx() sctx.ServiceContext {
	migrationPath, err := filepath.Abs("./services/location/migrations")
	if err != nil {
		migrationPath = "./services/location/migrations"
	}

	return sctx.NewServiceContext(
		sctx.WithName("Location service"),
		sctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(core.KeyCompPostgres, "", migrationPath)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(core.KeyCompGIN).(core.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), middleware.Recovery(serviceCtx))

		v1 := router.Group("/api/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	countryApiTransport := composer.NewComposerCountryApiTransport(serviceCtx)
	provinceApiTransport := composer.NewComposerProvinceApiTransport(serviceCtx)
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
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
