package custom_errors

import "errors"

var (
	ErrAlreadyRegistered = errors.New("such a user is already registered")
	ErrTokenOrUserInvalid = errors.New("this combination of login and token is not valid")
)