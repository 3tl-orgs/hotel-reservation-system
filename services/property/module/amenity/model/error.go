package model

import "errors"

var (
	ErrAmenityNameIsEmpty  = errors.New("amenity name is empty")
	ErrCannotCreateAmenity = errors.New("cannot create amenity")
	ErrCannotGetAmenity    = errors.New("cannot get amenity")
)
