package auth

import "errors"

var (
	ErrUserNotFound   = errors.New("err: User not found")
	ErrBadCredentials = errors.New("err: Bad Credentials")
)
