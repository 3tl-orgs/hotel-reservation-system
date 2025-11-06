package propertymodel

import "errors"

var (
	ErrAddressIsEmpty        = errors.New("address is empty")
	ErrPropertyTypeIdIsEmpty = errors.New("property type id is empty")
	ErrNameIsEmpty           = errors.New("property name is empty")
	ErrWarIdIsEmpty          = errors.New("war id is empty")
	ErrProvinceIdIsEmpty     = errors.New("province id is empty")
	ErrCountryIdIsEmpty      = errors.New("country id is empty")
	ErrCannotCreateProperty  = errors.New("cannot create property")
)
