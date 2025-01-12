package mock

import "errors"

type ResponseError struct {
	Message string `json:"message"`
}

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("your requested Item is not found")
	ErrConflict            = errors.New("your Item already exist")
)
