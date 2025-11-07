package model

import "errors"

var (
	PropertyTypeNameNotEmptyErr    = errors.New("property name not empty")
	PropertyTypeNameIsDuplicateErr = errors.New("property name is duplicated")
	CannotCreatePropertyTypeErr    = errors.New("cannot create property type")
	CannotGetPropertyTypeErr       = errors.New("cannot get property type")
)
