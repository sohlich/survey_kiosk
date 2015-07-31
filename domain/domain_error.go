package domain

import ()

type ValidationError struct {
	InternalError error
}

func (err *ValidationError) Error() string {
	return "Object validation failed"
}

type OperationFailError struct {
	InternalError error
}

func (err *OperationFailError) Error() string {
	return "Database operation failed!"
}
