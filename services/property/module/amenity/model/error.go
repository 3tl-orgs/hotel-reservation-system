package model

import "errors"

var (
	ErrAmenityNameIsEmpty  = errors.New("amenity name is empty")
	ErrCannotCreateAmenity = errors.New("cannot create amenity")
	ErrCannotGetAmenity    = errors.New("cannot get amenity")
	ErrCannotUpdateAmenity = errors.New("cannot update amenity")
	ErrAmenityNotFound = errors.New("cannot found amenity")
	ErrCannotDeleteAmenity = errors.New("cannot delete amenity")
	ErrCannotDeleteManyAmenities = errors.New("cannot delete many amenities")
)
