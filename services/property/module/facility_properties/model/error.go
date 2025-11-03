package facilitypropertiesmodel

import "github.com/pkg/errors"

var (
	ErrFacilityPropertiesNameIsEmpty  = errors.New("FacilityProperties name is empty")
	ErrCannotCreateFacilityProperties = errors.New("cannot create FacilityProperties")
	ErrCannotGetFacilityProperties    = errors.New("cannot get FacilityProperties")
	ErrCannotUpdateFacilityProperties = errors.New("cannot update FacilityProperties")
	ErrFacilityPropertiesNotFound     = errors.New("cannot found FacilityProperties")
	ErrCannotDeleteFacilityProperties = errors.New("cannot delete FacilityProperties")
	ErrRecordHasBeenDeleted           = errors.New("record has been deleted")
	ErrCannotGetProperty              = errors.New("cannot get property")
)
