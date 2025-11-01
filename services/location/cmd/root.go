package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/composer"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/ginc/middleware"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/gormc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func newServiceCtx() sctx.ServiceContext {
	migrationPath, err := filepath.Abs("./services/location/migrations")
	if err != nil {
		migrationPath = "./services/location/migrations"
	}
	log.Println("migrationPath:", migrationPath)
	return sctx.NewServiceContext(
		sctx.WithName("Location service"),
		sctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(core.KeyCompPostgres, "", migrationPath)),
		sctx.WithComponent(NewGrpcConfig()),
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

		go StartGRPCServices(serviceCtx)
		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

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

func StartGRPCServices(serviceCtx sctx.ServiceContext) {
	configComp := serviceCtx.MustGet(core.KeyCompGrpcConf).(core.GrpcConfig)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterCountryServiceServer(s, composer.NewCountryGrpcTransport(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
