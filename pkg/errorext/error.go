package errorext

import (
	"errors"
	"net/http"
)

func BuildCustomError(err error) error {
	var customErr CustomError
	ok := errors.As(err, &customErr)
	if ok {
		return &customErr
	}
	// return custom error with code
	return NewCustomError(http.StatusInternalServerError, err)
}

func ParseCustomError(err error) *CustomError {
	var customErr CustomError
	ok := errors.As(err, &customErr)
	if ok {
		return &customErr
	}
	// return custom error with code
	return NewCustomError(http.StatusInternalServerError, err)
}
