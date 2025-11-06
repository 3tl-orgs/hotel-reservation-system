package model

import "errors"

var (
	ErrFacilityNameEmpty = errors.New("facility name is empty")
	ErrInvalidIcon = errors.New("invalid icon")
	ErrCannotUploadIcon =  errors.New("cannot upload icon to cloud")

	ErrCannotCreateFacility = errors.New("cannot create facility")
	ErrCannotReadFacility = errors.New("cannot read facility")
	ErrCannotUpdateFacility = errors.New("cannot update facility")
	ErrCannotDeleteFacility = errors.New("cannot delete facility")
)