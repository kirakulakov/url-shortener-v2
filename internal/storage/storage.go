package storage

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLNotExist = errors.New("URL not exist")
)
