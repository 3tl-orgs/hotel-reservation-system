package roomtypemodel

import "errors"

var (
	ErrRoomNameIsEmpty      = errors.New("room name is empty")
	ErrRoomCodeIsEmpty      = errors.New("room code is empty")
	ErrRoomCodeIsDuplicated = errors.New("room code is duplicated")
	ErrCannotCreateRoom     = errors.New("cannot create room")
	ErrCannotGetRoom        = errors.New("cannot get room")
	ErrRoomNotFound         = errors.New("room not found")
	ErrRoomHasBeenDeleted   = errors.New("room already been deleted")
	ErrCannotUpdateRoom     = errors.New("cannot update room")
	ErrCannotDeleteRoom     = errors.New("cannot delete room")
)
