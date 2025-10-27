package wardsmodel

import "errors"

var (
	ErrWardNameIsEmpty      = errors.New("ward name is empty")
	ErrWardCodeIsEmpty      = errors.New("ward code is empty")
	ErrWardCodeIsDuplicated = errors.New("ward code is duplicated")
	ErrCannotCreateWard     = errors.New("cannot create ward")
	ErrCannotGetWard        = errors.New("cannot get ward")
	ErrWardNotFound         = errors.New("ward not found")
	ErrExistingWard         = errors.New("ward already exists")
	ErrCannotUpdateWard     = errors.New("cannot update ward")
	ErrCannotDeleteWard     = errors.New("cannot delete ward")
	ErrWardHasBeenDeleted   = errors.New("ward has been deleted")
	ErrIdDuplicated         = errors.New("id already exists")
)
