package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo/postgres"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	facilitypropertiesbiz "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/biz"
	facilitypropertiesrepo "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/repo"
	facilitypropertiesapi "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/transport/api"
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

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := postgres.NewPostgresRepo(db.GetDB())

	locationClient := sctx.MustGet(core.KeyCountryCompLocationClient).(*grpcclient.CountryRPCComponent)

	amenityService := biz.NewBusiness(amenityRepo, locationClient)
	amenityApi := api.NewAmenityApi(amenityService)
	return amenityApi
}

type FacilityPropertiesApiTransport interface {
	CreateFacilityPropHdl() gin.HandlerFunc
	UpdateFacilityPropHdl() gin.HandlerFunc
	GetFacilityPropByIdHdl() gin.HandlerFunc
	DeleteFacilityPropHdl() gin.HandlerFunc
	ListFacilityPropHdl() gin.HandlerFunc
	GetFacilityByPropHdl() gin.HandlerFunc
}

func ComposerFacilityPropertiesApiTransport(sctx srvctx.ServiceContext) FacilityPropertiesApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	facilityPropertiesRepo := facilitypropertiesrepo.NewFacilityPropertiesRepo(db.GetDB())
	facilityPropertiesBiz := facilitypropertiesbiz.NewFacilityPropertiesRepo(facilityPropertiesRepo)
	facilityPropertiesTransport := facilitypropertiesapi.NewFacilityPropertiesTransport(facilityPropertiesBiz)
	return facilityPropertiesTransport
}
