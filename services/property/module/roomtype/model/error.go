package roomtypemodel

import "errors"

var (
	ErrRoomNameIsEmpty             = errors.New("room name is empty")
	ErrInputData                   = errors.New("your input not have enough data")
	ErrRoomIsDuplicated            = errors.New("room is duplicated")
	ErrCannotCreateRoom            = errors.New("cannot create room")
	ErrCannotGetRoom               = errors.New("cannot get room")
	ErrRoomNotFound                = errors.New("room not found")
	ErrRoomHasBeenDeleted          = errors.New("room already been deleted")
	ErrCannotUpdateRoom            = errors.New("cannot update room")
	ErrCannotDeleteRoom            = errors.New("cannot delete room")
	ErrYourPropertyHaveNoRoomTypes = errors.New("you have no room types")
)
