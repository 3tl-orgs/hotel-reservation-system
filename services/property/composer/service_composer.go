package composer

import (
	"os"

	"github.com/gin-gonic/gin"
	amenitybiz "github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	amenityrepo "github.com/ngleanhvu/go-booking/services/property/module/amenity/repo"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	facilitybiz "github.com/ngleanhvu/go-booking/services/property/module/facility/biz"
	facilityrepo "github.com/ngleanhvu/go-booking/services/property/module/facility/repo"
	facilityapi "github.com/ngleanhvu/go-booking/services/property/module/facility/transport/api"
	facilitypropertiesbiz "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/biz"
	facilitypropertiesrepo "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/repo"
	facilitypropertiesapi "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/transport/api"
	propertybiz "github.com/ngleanhvu/go-booking/services/property/module/property/biz"
	propertystore "github.com/ngleanhvu/go-booking/services/property/module/property/repo"
	propertyapi "github.com/ngleanhvu/go-booking/services/property/module/property/transport/api"
	propertytypebiz "github.com/ngleanhvu/go-booking/services/property/module/propertytype/biz"
	propertytyperepo "github.com/ngleanhvu/go-booking/services/property/module/propertytype/repo"
	propertytypeapi "github.com/ngleanhvu/go-booking/services/property/module/propertytype/transport/api"
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
}

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

type FacilityPropertiesApiTransport interface {
	CreateFacilityPropHdl() gin.HandlerFunc
	UpdateFacilityPropHdl() gin.HandlerFunc
	GetFacilityPropByIdHdl() gin.HandlerFunc
	DeleteFacilityPropHdl() gin.HandlerFunc
	ListFacilityPropHdl() gin.HandlerFunc
	GetFacilityByPropHdl() gin.HandlerFunc
}

type PropertyApiTransport interface {
	CreatePropertyHdl() gin.HandlerFunc
}

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := amenityrepo.NewPostgresRepo(db.GetDB())

	countryClient := sctx.MustGet(core.KeyCountryCompLocationClient).(*grpcclient.CountryRPCComponent)

	amenityService := amenitybiz.NewBusiness(amenityRepo, countryClient)
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
	propertyTypeRepo := propertytyperepo.NewPostgresRepo(db.GetDB())
	propertyTypeService := propertytypebiz.NewBusiness(propertyTypeRepo)
	propertyTypeApi := propertytypeapi.NewPropertyTypeApi(propertyTypeService)
	return propertyTypeApi
}

func ComposerRoomTypeApiTransport(sctx srvctx.ServiceContext) RoomTypeApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	roomTypeRepo := roomtyperepo.NewRoomTypeRepo(db.GetDB())
	roomTypeBiz := roomtypebiz.NewRoomTypeBusiness(roomTypeRepo)
	roomTypeTransport := roomtypeapi.NewRoomTypeTransport(roomTypeBiz)
	return roomTypeTransport
}

func ComposerPropertyTransport(sctx srvctx.ServiceContext) PropertyApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	countryClient := sctx.MustGet(core.KeyCountryCompLocationClient).(*grpcclient.CountryRPCComponent)
	provinceClient := sctx.MustGet(core.KeyProvinceCompLocationClient).(*grpcclient.ProvinceRPCComponent)
	wardClient := sctx.MustGet(core.KeyWardCompLocationClient).(*grpcclient.WardRPCComponent)
	propertyRepo := propertystore.NewPropertyStore(db.GetDB())
	propertyTypeRepo := propertytyperepo.NewPostgresRepo(db.GetDB())
	bucketName := os.Getenv("R2_BUCKET_NAME")
	accessKey := os.Getenv("R2_ACCESS_KEY_ID")
	secretKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	accountId := os.Getenv("R2_ACCOUNT_ID")
	domain := os.Getenv("R2_PUBLIC_DOMAIN")
	s3Uploader := uploadprovider.NewR2Provider(bucketName, accountId, accessKey, secretKey, domain)
	propertyBiz := propertybiz.NewPropertyBiz(propertyRepo, s3Uploader, propertyTypeRepo, countryClient, provinceClient, wardClient)
	propertyApi := propertyapi.NewPropertyApi(propertyBiz)
	return propertyApi
}
