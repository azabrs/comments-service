package custom_errors

import "errors"

var (
	ErrAlreadyRegistered = errors.New("such a user is already registered")

)