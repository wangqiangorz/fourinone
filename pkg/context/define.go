package context

import "errors"

const (
	WORKERPRFIX = "__worker__"
)

var (
	ERRORCONFIGSETTING  = errors.New("config setting is wrong!")
	ERRORGETGROUPSERVER = errors.New("The length of groupserver is 0")
)
