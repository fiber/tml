package tml

import (
	"errors"
)

var (
	ErrAbstractCall = errors.New("call to abstract function")
	ErrInvalid      = errors.New("invalid request or function call")
)
