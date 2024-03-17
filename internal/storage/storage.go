package storage

import "errors"

var (
	ErrURLNotFound  = errors.New("URL not found")
	ErrURLNotExists = errors.New("URL not exists")
	ErrURLExists    = errors.New("URL already exists")
)
