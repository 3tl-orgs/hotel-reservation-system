package provincemodel

import "errors"

var (
	ErrProvinceNameIsEmpty      = errors.New("country name is empty")
	ErrProvinceCodeIsEmpty      = errors.New("country code is empty")
	ErrProvinceCodeIsDuplicated = errors.New("country code is duplicated")
	ErrCannotCreateProvince     = errors.New("cannot create country")
	ErrCannotGetProvince        = errors.New("cannot get country")
	ErrProvinceNotFound         = errors.New("country not found")
	ErrExistingProvince         = errors.New("country already exists")
	ErrCannotUpdateProvince     = errors.New("cannot update country")
	ErrCannotDeleteProvince     = errors.New("cannot delete country")
)
