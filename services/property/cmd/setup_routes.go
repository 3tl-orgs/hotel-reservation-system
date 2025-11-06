package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/ngleanhvu/go-booking/services/property/composer"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	amenityApiTransport := composer.ComposerAmenityApiTransport(serviceCtx)
	propertyTypeApiTransport := composer.ComposerPropertTypeApiTransport(serviceCtx)
	facilityPropertiesApiTransport := composer.ComposerFacilityPropertiesApiTransport(serviceCtx)
	propertyApiTransport := composer.ComposerPropertyApiTransport(serviceCtx)

	properties := router.Group("/properties")
	{
		// amenity
		// Amenities API
		properties.POST("/amenities", amenityApiTransport.CreateAmenityHdl())
		properties.GET("/amenities/:id", amenityApiTransport.GetAmenityByIdHdl())
		properties.PATCH("/amenities/:id", amenityApiTransport.UpdateAmenityHdl())
		properties.DELETE("/amenities/:id", amenityApiTransport.DeleteAmenityByIdHdl())
		properties.GET("/amenities/test-grpc/:id", amenityApiTransport.TestGrpcHdl())

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

		// Property API
		properties.POST("/property", propertyApiTransport.CreatePropertyHdl())
		properties.GET("/property/:id", propertyApiTransport.GetPropertyByIdHdl())
		properties.PATCH("/property/:id", propertyApiTransport.UpdatePropertyHdl())
		properties.DELETE("/property/:id", propertyApiTransport.DeletePropertyHdl())
		properties.GET("/property", propertyApiTransport.ListPropertyHdl())
	}

}
