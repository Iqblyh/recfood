package errorConv

import (
	"errors"
	"strings"
)

func Conversion(err error) error {
	var errNew error
	if strings.Contains(err.Error(), "not found") {
		errNew = errors.New(ErrDBNotFound)
	}

	if strings.Contains(err.Error(), "invalid") {
		errNew = errors.New(ErrInvalid)
	}

	if strings.Contains(err.Error(), "mismatch") {
		errNew = errors.New(ErrMismatch)
	}

	return errNew
}
