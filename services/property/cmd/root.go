package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/composer"
	"github.com/ngleanhvu/go-booking/shared/core"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc/middleware"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/gormc"
	"github.com/spf13/cobra"
)

func newServiceCtx() sctx.ServiceContext {
	// migrationPath, err := filepath.Abs("./services/property/migrations")
	// if err != nil {
	// 	migrationPath = "./services/property/migrations"
	// }

	// return sctx.NewServiceContext(
	// 	sctx.WithName("Property service"),
	// 	sctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
	// 	sctx.WithComponent(gormc.NewGormDB(core.KeyCompPostgres, "", migrationPath)),
	// )

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../migrations")

	log.Println("basePath:", basePath)

	return sctx.NewServiceContext(
		sctx.WithName("Property service"),
		sctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(core.KeyCompPostgres, "", basePath)),
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
	properties := router.Group("/properties")
	{
		// amenity
		properties.POST("/amenities", amenityApiTransport.CreateAmenityHdl())
		properties.GET("/amenities/:id", amenityApiTransport.GetAmenityByIdHdl())
		properties.PATCH("/amenities/:id", amenityApiTransport.UpdateAmenityHdl())
		properties.DELETE("/amenities/:id", amenityApiTransport.DeleteAmenityByIdHdl())

		//Get by ids, delete by ids -> list ids[] in body
		// properties.DELETE("/amenities/bulk", amenityApiTransport.DeleteAmenityByIdsHdl())
		// properties.GET("/amenities/bulk", amenityApiTransport.GetAmenityByIdsHdl())

	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
