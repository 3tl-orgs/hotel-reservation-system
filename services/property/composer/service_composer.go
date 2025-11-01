package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/composer/client"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo/postgres"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo/rpc/location"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
)

type AmenityApiTransport interface {
	CreateAmenityHdl() gin.HandlerFunc
	GetAmenityByIdsHdl() gin.HandlerFunc
	UpdateAmenityHdl() gin.HandlerFunc
	DeleteAmenityByIdHdl() gin.HandlerFunc
	GetAmenityByIdHdl() gin.HandlerFunc
	DeleteAmenityByIdsHdl() gin.HandlerFunc
	TestGrpcHdl() gin.HandlerFunc
}

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := postgres.NewPostgresRepo(db.GetDB())
	countryRepo := location.NewCountryClient(client.ComposerCountryRpcClient(sctx))
	amenityService := biz.NewBusiness(amenityRepo, countryRepo)
	amenityApi := api.NewAmenityApi(amenityService)
	return amenityApi
}
