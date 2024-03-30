package domain

import "errors"

var (
	ErrorProductNotFound = errors.New("error product not found")
	ErrProductDoesNotExist = errors.New("error product does not exist")
)