package customErrors

import "github.com/pkg/errors"

var (
	ErrVersionNotFound = errors.New("Version not found")
)
