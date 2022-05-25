package resp

import "errors"

var (
	Unauthorized    = errors.New("unauthorized request")
	InvalidPassword = errors.New("invalid password")
	InvalidParam    = errors.New("invalid request param")
	InvalidProject  = errors.New("duplicate project name")
)
