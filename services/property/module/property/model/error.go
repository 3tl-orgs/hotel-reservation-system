package propertymodel

import "github.com/pkg/errors"

var (
	ErrIsEmpty                  = errors.New("Data is empty")
	ErrCannotCreateProperty     = errors.New("cannot create Property")
	ErrCannotGetProperty        = errors.New("cannot get Property")
	ErrCannotUpdateProperty     = errors.New("cannot update Property")
	ErrPropertyNotFound         = errors.New("cannot found Property")
	ErrCannotDeleteProperty     = errors.New("cannot delete Property")
	ErrRecordHasBeenDeleted     = errors.New("record has been deleted")
	ErrIdOrAmenitiesNotAccepted = errors.New("Id or amenities not accepted")
)
