package core

import (
	"errors"
)

// ErrInvalidValue signals that a nil value has been provided
var ErrInvalidValue = errors.New("invalid value provided")