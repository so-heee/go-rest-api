package repository

import (
	"errors"
)

var (
	ErrNilID         = errors.New("nil id")
	ErrInvalidID     = errors.New("invalid uuid")
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
	ErrAlreadyExists = errors.New("already exists")
	ErrInvalidArg    = errors.New("argument error")
	ErrBind          = errors.New("bind error")
	ErrValidate      = errors.New("validate error")
)
