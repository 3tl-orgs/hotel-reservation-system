package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/composer"
	"github.com/ngleanhvu/go-booking/shared/core"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc/middleware"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/gormc"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/grpcclient"
	"github.com/spf13/cobra"
)

func newServiceCtx() sctx.ServiceContext {
	migrationPath, err := filepath.Abs("./services/property/migrations")
	if err != nil {
		migrationPath = "./services/property/migrations"
	}

	return sctx.NewServiceContext(
		sctx.WithName("Property service"),
		sctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(core.KeyCompPostgres, "", migrationPath)),
		sctx.WithComponent(NewGrpcConfig()),
		sctx.WithComponent(grpcclient.NewCountryRPCClientComponent(core.KeyCountryCompLocationClient, "")),
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
	amenityApiTransport := composer.ComposerAmenityApiTransport(serviceCtx)
	propertyTypeApiTransport := composer.ComposerPropertTypeApiTransport(serviceCtx)
	facilityPropertiesApiTransport := composer.ComposerFacilityPropertiesApiTransport(serviceCtx)
	roomTypeApiTransport := composer.ComposerRoomTypeApiTransport(serviceCtx)

	properties := router.Group("/properties")
	{
		// amenity
		// Amenities API
		properties.POST("/amenities", amenityApiTransport.CreateAmenityHdl())
		properties.GET("/amenities/:id", amenityApiTransport.GetAmenityByIdHdl())
		properties.PATCH("/amenities/:id", amenityApiTransport.UpdateAmenityHdl())
		properties.DELETE("/amenities/:id", amenityApiTransport.DeleteAmenityByIdHdl())
		properties.GET("/amenities/test-grpc/:id", amenityApiTransport.TestGrpcHdl())

		// property type
		properties.GET("/property-types/:id", propertyTypeApiTransport.GetPropertyTypeByIdHdl())
		properties.POST("/property-types", propertyTypeApiTransport.CreatePropertyTypeHdl())
		properties.GET("/property-types/all", propertyTypeApiTransport.ListPropertyTypeHdl())

		// Facility_Properties API
		properties.POST("/facility-properties", facilityPropertiesApiTransport.CreateFacilityPropHdl())
		properties.PATCH("/facility-properties/:id", facilityPropertiesApiTransport.UpdateFacilityPropHdl())
		properties.DELETE("/facility-properties/:id", facilityPropertiesApiTransport.DeleteFacilityPropHdl())
		properties.GET("/facility-properties/:id", facilityPropertiesApiTransport.GetFacilityPropByIdHdl())
		properties.GET("/facility-properties", facilityPropertiesApiTransport.ListFacilityPropHdl())
		properties.GET("/facility-properties/facilities/:id", facilityPropertiesApiTransport.GetFacilityByPropHdl())

		// Room Type API
		properties.POST("/roomtypes", roomTypeApiTransport.CreateRoomTypeHdl())
		properties.GET("/roomtypes/:id", roomTypeApiTransport.GetRoomTypeByIdHdl())
		properties.PATCH("/roomtypes/:id", roomTypeApiTransport.UpdateRoomTypeHdl())
		properties.DELETE("/roomtypes/:id", roomTypeApiTransport.DeleteRoomTypeHdl())
		properties.GET("/roomtypes", roomTypeApiTransport.ListRoomTypeHdl())
	}

}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
