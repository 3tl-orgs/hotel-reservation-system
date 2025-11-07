package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/composer"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	amenityApiTransport := composer.ComposerAmenityApiTransport(serviceCtx)
	facilityApiTransport := composer.ComposerFacilityApiTransport(serviceCtx)
	propertyTypeApiTransport := composer.ComposerPropertTypeApiTransport(serviceCtx)
	facilityPropertiesApiTransport := composer.ComposerFacilityPropertiesApiTransport(serviceCtx)
	roomTypeApiTransport := composer.ComposerRoomTypeApiTransport(serviceCtx)

	properties := router.Group("/properties")
	{
		// amenity
		// Amenities API
		properties.POST("/amenities", amenityApiTransport.CreateAmenityHdl())
		properties.GET("/amenities/:id", amenityApiTransport.GetAmenityByIdHdl())
		properties.PATCH("/amenities/:id", amenityApiTransport.UpdateAmenityHdl())
		properties.DELETE("/amenities/:id", amenityApiTransport.DeleteAmenityByIdHdl())
		properties.GET("/amenities/test-grpc/:id", amenityApiTransport.TestGrpcHdl())

		properties.POST("/facilities", facilityApiTransport.CreateFacilityHdl())
		properties.GET("/facilities", facilityApiTransport.GetListFacilitiesHdl())
		properties.GET("/facilities/:id", facilityApiTransport.GetFacilityByIdHdl())
		properties.PATCH("/facilities/:id", facilityApiTransport.UpdateFacilityHdl())
		properties.DELETE("/facilities/:id", facilityApiTransport.DeleteFacilityHdl())

		// property type
		properties.GET("/property-types/:id", propertyTypeApiTransport.GetPropertyTypeByIdHdl())
		properties.POST("/property-types", propertyTypeApiTransport.CreatePropertyTypeHdl())
		properties.GET("/property-types/all", propertyTypeApiTransport.ListPropertyTypeHdl())

		// Facility_Properties API
		properties.POST("/facility-properties", facilityPropertiesApiTransport.CreateFacilityPropHdl())
		properties.PATCH("/facility-properties/:id", facilityPropertiesApiTransport.UpdateFacilityPropHdl())
		properties.DELETE("/facility-properties/:id", facilityPropertiesApiTransport.DeleteFacilityPropHdl())
		properties.GET("/facility-properties/:id", facilityPropertiesApiTransport.GetFacilityPropByIdHdl())
		properties.GET("/facility-properties", facilityPropertiesApiTransport.ListFacilityPropHdl())
		properties.GET("/facility-properties/facilities/:id", facilityPropertiesApiTransport.GetFacilityByPropHdl())

		// Room Type API
		properties.POST("/roomtypes", roomTypeApiTransport.CreateRoomTypeHdl())
		properties.GET("/roomtypes/:id", roomTypeApiTransport.GetRoomTypeByIdHdl())
		properties.PATCH("/roomtypes/:id", roomTypeApiTransport.UpdateRoomTypeHdl())
		properties.DELETE("/roomtypes/:id", roomTypeApiTransport.DeleteRoomTypeHdl())
		properties.GET("/roomtypes", roomTypeApiTransport.ListRoomTypeHdl())
		properties.GET("/roomtypes/property/:id", roomTypeApiTransport.GetRoomTypeByPropHdl())
	}

}
