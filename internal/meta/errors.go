package meta

import "errors"

var (
	ErrNotExist = errors.New("meta does not exist")
	ErrInvalid  = errors.New("invalid meta")
)
