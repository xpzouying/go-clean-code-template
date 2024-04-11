package api

import "errors"

var (
	ErrEmptyUsername = errors.New("empty username")
	ErrEmptyAvatar   = errors.New("empty avatar")
)
