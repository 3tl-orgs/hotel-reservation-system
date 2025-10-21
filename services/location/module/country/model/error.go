package model

import "errors"

var (
	ErrCountryNameIsEmpty      = errors.New("country name is empty")
	ErrCountryCodeIsEmpty      = errors.New("country code is empty")
	ErrCountryCodeIsDuplicated = errors.New("country code is duplicated")
	ErrCannotCreateCountry     = errors.New("cannot create country")
)
