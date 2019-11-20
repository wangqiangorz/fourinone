package service

import "errors"

var (
	ERRORGETLEADER = errors.New("All server can not connnect")
	KEYNOTEXIST    = errors.New("domainNodeKey is not exist")
	KEYEXIST       = errors.New("domainNodeKey is already exist")
)
