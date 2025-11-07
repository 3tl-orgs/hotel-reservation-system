package composer

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo/postgres"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	facilitybiz "github.com/ngleanhvu/go-booking/services/property/module/facility/biz"
	facilityrepo "github.com/ngleanhvu/go-booking/services/property/module/facility/repo"
	facilityapi "github.com/ngleanhvu/go-booking/services/property/module/facility/transport/api"
	facilitypropertiesbiz "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/biz"
	facilitypropertiesrepo "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/repo"
	facilitypropertiesapi "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/transport/api"
	propertyTypeBiz "github.com/ngleanhvu/go-booking/services/property/module/propertytype/biz"
	propertyTypeRepo1 "github.com/ngleanhvu/go-booking/services/property/module/propertytype/repo"
	propertyTypeTransport "github.com/ngleanhvu/go-booking/services/property/module/propertytype/transport/api"
	roomtypebiz "github.com/ngleanhvu/go-booking/services/property/module/roomtype/biz"
	roomtyperepo "github.com/ngleanhvu/go-booking/services/property/module/roomtype/repo"
	roomtypeapi "github.com/ngleanhvu/go-booking/services/property/module/roomtype/transport/api"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/ngleanhvu/go-booking/shared/srvctx/component/grpcclient"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
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

type FacilityApiTransport interface {
	CreateFacilityHdl() gin.HandlerFunc
	GetFacilityByIdHdl() gin.HandlerFunc
	GetListFacilitiesHdl() gin.HandlerFunc
	UpdateFacilityHdl() gin.HandlerFunc
	DeleteFacilityHdl() gin.HandlerFunc
type PropertyTypeApiTransport interface {
	CreatePropertyTypeHdl() gin.HandlerFunc
	GetPropertyTypeByIdHdl() gin.HandlerFunc
	ListPropertyTypeHdl() gin.HandlerFunc
}

type RoomTypeApiTransport interface {
	CreateRoomTypeHdl() gin.HandlerFunc
	DeleteRoomTypeHdl() gin.HandlerFunc
	UpdateRoomTypeHdl() gin.HandlerFunc
	GetRoomTypeByIdHdl() gin.HandlerFunc
	ListRoomTypeHdl() gin.HandlerFunc
	GetRoomTypeByPropHdl() gin.HandlerFunc
}

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := postgres.NewPostgresRepo(db.GetDB())

	locationClient := sctx.MustGet(core.KeyCountryCompLocationClient).(*grpcclient.CountryRPCComponent)

	amenityService := biz.NewBusiness(amenityRepo, locationClient)
	amenityApi := api.NewAmenityApi(amenityService)
	return amenityApi
}

func ComposerFacilityApiTransport(srvctx srvctx.ServiceContext) FacilityApiTransport {
	db := srvctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	facilityRepo := facilityrepo.NewFacilityRepo(db.GetDB())

	bucketName := os.Getenv("R2_BUCKET_NAME")
	accessKey := os.Getenv("R2_ACCESS_KEY_ID")
	secretKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	accountId := os.Getenv("R2_ACCOUNT_ID")
	domain := os.Getenv("R2_PUBLIC_DOMAIN")

	s3Uploader := uploadprovider.NewR2Provider(bucketName, accountId, accessKey, secretKey, domain)
	facilityBiz := facilitybiz.NewFacilityBusiness(facilityRepo, s3Uploader)

	facilityApi := facilityapi.NewFacilityApi(facilityBiz)
	return facilityApi
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

func ComposerPropertTypeApiTransport(sctx srvctx.ServiceContext) PropertyTypeApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	propertyTypeRepo := propertyTypeRepo1.NewPostgresRepo(db.GetDB())
	propertyTypeService := propertyTypeBiz.NewBusiness(propertyTypeRepo)
	propertyTypeApi := propertyTypeTransport.NewPropertyTypeApi(propertyTypeService)
	return propertyTypeApi
}

func ComposerRoomTypeApiTransport(sctx srvctx.ServiceContext) RoomTypeApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	roomTypeRepo := roomtyperepo.NewRoomTypeRepo(db.GetDB())
	roomTypeBiz := roomtypebiz.NewRoomTypeBusiness(roomTypeRepo)
	roomTypeTransport := roomtypeapi.NewRoomTypeTransport(roomTypeBiz)
	return roomTypeTransport
}
