package composer

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/biz"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/repo"
	"github.com/ngleanhvu/go-booking/services/property/module/amenity/transport/api"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
)

type AmenityApiTransport interface {
	CreateAmenityHdl() gin.HandlerFunc
	GetAmenityByIdsHdl() gin.HandlerFunc
}

func ComposerAmenityApiTransport(sctx srvctx.ServiceContext) AmenityApiTransport {
	db := sctx.MustGet(core.KeyCompPostgres).(core.GormComponent)
	amenityRepo := repo.NewPostgresRepo(db.GetDB())
	amenityService := biz.NewBusiness(amenityRepo)
	amenityApi := api.NewAmenityApi(amenityService)
	return amenityApi
}
