package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrUserNotFound struct {
	Err error
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrUserNotFound() *ErrUserNotFound {
	return &ErrUserNotFound{Err: errors.New("user not found")}
}

type ErrAlreadyExists struct {
	Err error
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrAlreadyExists() *ErrAlreadyExists {
	return &ErrAlreadyExists{Err: errors.New("user already exists")}
}

type ErrUnkown struct {
	Err error
}

func (e *ErrUnkown) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrUnkown() *ErrUnkown {
	return &ErrUnkown{Err: errors.New("unkown error")}
}

type ErrInvalidArguments struct {
	Err error
}

func (e *ErrInvalidArguments) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrInvalidArguments() *ErrInvalidArguments {
	return &ErrInvalidArguments{Err: errors.New("invalid argument(s)")}
}

type ErrInvalidCredentials struct {
	Err error
}

func (e *ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewErrInvalidCredentials() *ErrInvalidCredentials {
	return &ErrInvalidCredentials{Err: errors.New("invalid credential(s)")}
}

// Receives a grpc code, returns the corresponding custom error
// Used in http service - repository layer
func ErrFromGRPCcode(code int32) error {
	var err error
	switch code {
	case 5:
		err = NewErrUserNotFound()
	case 6:
		err = NewErrAlreadyExists()
	case 3:
		err = NewErrInvalidArguments()
	case 7:
		err = NewErrInvalidCredentials()
	default:
		err = NewErrUnkown()
	}
	return err
}

// Selects the correct http status regarding the error caught
// Used in http service - transport layer
func ErrToHTTPStatus(err error) int {
	switch err.(type) {
	case *ErrUserNotFound:
		return http.StatusNotFound
	case *ErrAlreadyExists:
		return http.StatusConflict
	case *ErrInvalidArguments:
		return http.StatusBadRequest
	case *ErrInvalidCredentials:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
