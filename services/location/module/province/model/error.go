package provincemodel

import "errors"

var (
	ErrProvinceNameIsEmpty      = errors.New("province name is empty")
	ErrProvinceCodeIsEmpty      = errors.New("province code is empty")
	ErrProvinceCodeIsDuplicated = errors.New("province code is duplicated")
	ErrCannotCreateProvince     = errors.New("cannot create province")
	ErrCannotGetProvince        = errors.New("cannot get province")
	ErrProvinceNotFound         = errors.New("province not found")
	ErrExistingProvince         = errors.New("province already exists")
	ErrCannotUpdateProvince     = errors.New("cannot update province")
	ErrCannotDeleteProvince     = errors.New("cannot delete province")
	ErrProvinceHasBeenDeleted   = errors.New("province has been deleted")
	ErrIdDuplicated             = errors.New("id already exists")
)
