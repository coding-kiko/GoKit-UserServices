package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coding-kiko/GoKit-Project-Bootcamp/GRPCServiceA/pkg/user/proto"
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

// Receives a custom error, returns the corresponding proto status struct filled
// Used in grpc service - transport layer
func ErrToGRPCcode(e error) *proto.Status {
	var status proto.Status
	switch e.(type) {
	case *ErrUserNotFound:
		status.Code = 5
		status.Message = e.Error()
	case *ErrAlreadyExists:
		status.Code = 6
		status.Message = e.Error()
	default:
		status.Code = 2
		status.Message = "unkown error"
	}
	return &status
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
	default:
		return http.StatusInternalServerError
	}
}
