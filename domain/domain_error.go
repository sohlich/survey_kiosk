package domain

import ()

type ValidationError struct {
	InternalError error
}

func (err *ValidationError) Error() string {
	return "Object validation failed"
}
