package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/business"
	"github.com/ngleanhvu/go-booking/services/location/module/country/repo"
	"github.com/ngleanhvu/go-booking/services/location/module/country/transport/api"
	"github.com/ngleanhvu/go-booking/services/location/module/country/transport/rpc"
	provincebusiness "github.com/ngleanhvu/go-booking/services/location/module/province/business"
	provincerepo "github.com/ngleanhvu/go-booking/services/location/module/province/repo"
	provincetransport "github.com/ngleanhvu/go-booking/services/location/module/province/transport"
	wardsbusiness "github.com/ngleanhvu/go-booking/services/location/module/ward/business"
	wardsrepo "github.com/ngleanhvu/go-booking/services/location/module/ward/repo"
	wardstransport "github.com/ngleanhvu/go-booking/services/location/module/ward/transport"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
)

type CountryApiTransport interface {
	CreateCountryHdl() gin.HandlerFunc
	GetCountryByIdHdl() gin.HandlerFunc
	UpdateCountryHdl() gin.HandlerFunc
	ListCountryHdl() gin.HandlerFunc
	GetCountryByCodeHdl() gin.HandlerFunc
	DeleteCountryHdl() gin.HandlerFunc
}

type ProvinceApiTransport interface {
	CreateProvinceHdl() gin.HandlerFunc
	GetProvinceByIdHdl() gin.HandlerFunc
	UpdateProvinceHdl() gin.HandlerFunc
	GetProvinceByCodeHdl() gin.HandlerFunc
	DeleteProvinceHdl() gin.HandlerFunc
	ListProvinceHdl() gin.HandlerFunc
}

type WardApiTransport interface {
	CreateWardHdl() gin.HandlerFunc
	GetWardByIdHdl() gin.HandlerFunc
	UpdateWardHdl() gin.HandlerFunc
	DeleteWardHdl() gin.HandlerFunc
	GetWardByCodeHdl() gin.HandlerFunc
	ListWardHdl() gin.HandlerFunc
}

func NewComposerCountryApiTransport(sctx srvctx.ServiceContext) CountryApiTransport {
	countryDb := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	countryRepo := repo.NewPostgresRepo(countryDb.GetDB())
	countryBiz := business.NewBusiness(countryRepo)
	countryApi := api.NewApi(countryBiz)
	return countryApi
}

func NewComposerProvinceApiTransport(sctx srvctx.ServiceContext) ProvinceApiTransport {
	provinceDb := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	provinceRepo := provincerepo.NewProvinceRepo(provinceDb.GetDB())
	countryRepo := repo.NewPostgresRepo(provinceDb.GetDB())
	provinceBusiness := provincebusiness.NewProvinceBusiness(provinceRepo, countryRepo)
	provinceTransport := provincetransport.NewProvinceTransport(provinceBusiness)
	return provinceTransport
}

func NewComposerWardApiTransport(sctx srvctx.ServiceContext) WardApiTransport {
	wardDb := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	wardRepo := wardsrepo.NewWardRepo(wardDb.GetDB())
	provinceRepo := provincerepo.NewProvinceRepo(wardDb.GetDB())
	wardBusiness := wardsbusiness.NewWardBusiness(wardRepo, provinceRepo)
	wardTransport := wardstransport.NewWardTransport(wardBusiness)
	return wardTransport
}

func NewCountryGrpcTransport(sctx srvctx.ServiceContext) pb.CountryServiceServer {
	countryDb := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	countryRepo := repo.NewPostgresRepo(countryDb.GetDB())
	countryBiz := business.NewBusiness(countryRepo)
	countryRpc := rpc.NewRpc(countryBiz)
	return countryRpc
}
