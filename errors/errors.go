package custom_errors

import "errors"

var (
	ErrAlreadyRegistered = errors.New("such a user is already registered")
	ErrTokenOrUserInvalid = errors.New("this combination of login and token is not valid")
	ErrPostIDIncorrect = errors.New("PostId and parent comment PostID did not match")
	ErrPostNotAvaliableToComment = errors.New("the author has limited the possibility of commenting under this post")
	ErrReachecMaxSub = errors.New("it is impossible to subscribe because the maximum possible number of subscriptions has been reached")
	ErrParentIdIncorrect = errors.New("there is no comment with such an ParentID")
)