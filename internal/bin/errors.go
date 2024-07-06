package bin

import "errors"

var (
	ErrNotExist = errors.New("bin does not exist")
	ErrInvalid  = errors.New("invalid bin")
)
