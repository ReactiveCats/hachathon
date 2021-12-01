package platform

import (
	"errors"
	"fmt"
	"net/http"
)

// ServerError is like 5xx status
type ServerError struct {
	err error
}

func (i ServerError) Error() string {
	return i.err.Error()
}

func NewInternal(format string, args ...interface{}) error {
	return WrapInternal(fmt.Errorf(format, args...))
}

func WrapInternal(err error) error {
	return ServerError{
		err: err,
	}
}

// ClientError is like 4xx status
type ClientError interface {
	error
	Message() string
	HttpStatusCode() int
}

// ForbiddenError represents 403
type ForbiddenError struct {
	err     error
	message string
}

func Forbidden(message string) error {
	return WrapForbidden(errors.New(message), message)
}

func WrapForbidden(err error, message string) error {
	return ForbiddenError{
		err:     err,
		message: message,
	}
}

func (i ForbiddenError) Error() string {
	return i.err.Error()
}

func (i ForbiddenError) Message() string {
	return i.message
}

func (i ForbiddenError) HttpStatusCode() int {
	return http.StatusForbidden
}

// NotFoundError represents 404
type NotFoundError struct {
	err     error
	message string
}

func NotFound(message string) error {
	return NotFoundError{
		err:     errors.New(message),
		message: message,
	}
}

func (i NotFoundError) Error() string {
	return i.err.Error()
}

func (i NotFoundError) Message() string {
	return i.message
}

func (i NotFoundError) HttpStatusCode() int {
	return http.StatusNotFound
}

// ConflictError represents 409
type ConflictError struct {
	err     error
	message string
}

func Conflict(message string) error {
	return WrapConflict(errors.New(message), message)
}

func WrapConflict(err error, message string) error {
	return ConflictError{
		err:     err,
		message: message,
	}
}

func (i ConflictError) Error() string {
	return i.err.Error()
}

func (i ConflictError) Message() string {
	return i.message
}

func (i ConflictError) HttpStatusCode() int {
	return http.StatusConflict
}

// UnprocessableEntityError represents 422
type UnprocessableEntityError struct {
	err     error
	message string
}

func UnprocessableEntity(message string) error {
	return WrapUnprocessableEntity(errors.New(message), message)
}

func WrapUnprocessableEntity(err error, message string) error {
	return UnprocessableEntityError{
		err:     err,
		message: message,
	}
}

func (i UnprocessableEntityError) Error() string {
	return i.err.Error()
}

func (i UnprocessableEntityError) Message() string {
	return i.message
}

func (i UnprocessableEntityError) HttpStatusCode() int {
	return http.StatusUnprocessableEntity
}
