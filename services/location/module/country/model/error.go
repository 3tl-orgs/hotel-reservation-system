package model

import "errors"

var (
	ErrCountryNameIsEmpty      = errors.New("country name is empty")
	ErrCountryCodeIsEmpty      = errors.New("country code is empty")
	ErrCountryCodeIsDuplicated = errors.New("country code is duplicated")
	ErrCannotCreateCountry     = errors.New("cannot create country")
	ErrCannotGetCountry        = errors.New("cannot get country")
	ErrCountryNotFound         = errors.New("country not found")
	ErrExistingCountry         = errors.New("country already exists")
	ErrCannotUpdateCountry     = errors.New("cannot update country")
)
