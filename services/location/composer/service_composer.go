package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/location/module/country/business"
	"github.com/ngleanhvu/go-booking/services/location/module/country/repo"
	"github.com/ngleanhvu/go-booking/services/location/module/country/transport/api"
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

func NewComposerCountryApiTransport(sctx srvctx.ServiceContext) CountryApiTransport {
	countryDb := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	countryRepo := repo.NewPostgresRepo(countryDb.GetDB())
	countryBiz := business.NewBusiness(countryRepo)
	countryApi := api.NewApi(countryBiz)
	return countryApi
}
