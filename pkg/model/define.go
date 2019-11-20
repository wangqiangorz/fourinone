package model

import (
	"errors"
)

const (
	NOTREADY  = 1
	READY     = 0
	EXCEPTION = -1
)

var (
	ERRORPARSE = errors.New("parse error!")
)
