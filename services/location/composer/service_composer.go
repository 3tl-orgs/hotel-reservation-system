package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/business"
	"github.com/ngleanhvu/go-booking/services/location/module/country/repo"
	"github.com/ngleanhvu/go-booking/services/location/module/country/transport/api"
	provincebusiness "github.com/ngleanhvu/go-booking/services/location/module/province/business"
	provincerepo "github.com/ngleanhvu/go-booking/services/location/module/province/repo"
	provincetransport "github.com/ngleanhvu/go-booking/services/location/module/province/transport"
	"github.com/ngleanhvu/go-booking/shared/core"
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
	provinceBusiness := provincebusiness.NewProvinceBusiness(provinceRepo)
	provinceTransport := provincetransport.NewProvinceTransport(provinceBusiness)
	return provinceTransport
}
