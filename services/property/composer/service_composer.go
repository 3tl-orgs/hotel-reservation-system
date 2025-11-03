package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	propertyTypeBiz "github.com/ngleanhvu/go-booking/services/property/module/propertytype/biz"
	propertyTypeTransport "github.com/ngleanhvu/go-booking/services/property/module/propertytype/transport/api"
	propertyTypeRepo1 "github.com/ngleanhvu/go-booking/services/property/module/propertytype/repo"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo/postgres"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/grpcclient"
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

type PropertyTypeApiTransport interface {
	CreatePropertyTypeHdl() gin.HandlerFunc
	GetPropertyTypeByIdHdl() gin.HandlerFunc
}

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := postgres.NewPostgresRepo(db.GetDB())

	locationClient := sctx.MustGet(core.KeyCountryCompLocationClient).(*grpcclient.CountryRPCComponent)

	amenityService := biz.NewBusiness(amenityRepo, locationClient)
	amenityApi := api.NewAmenityApi(amenityService)
	return amenityApi
}

func ComposerPropertTypeApiTransport(sctx srvctx.ServiceContext) PropertyTypeApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	propertyTypeRepo := propertyTypeRepo1.NewPostgresRepo(db.GetDB())
	propertyTypeService := propertyTypeBiz.NewBusiness(propertyTypeRepo)
	propertyTypeApi := propertyTypeTransport.NewPropertyTypeApi(propertyTypeService)
	return propertyTypeApi
}
