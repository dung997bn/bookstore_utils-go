package resterrors

import (
	"errors"
	"net/http"
)

//RestErr type
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

//NewBadRequestError create new bad request status with message
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError create new not found status with message
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

//NewInternalServerError create new internal server error status with message
func NewInternalServerError(message string, err error) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Causes:  []interface{}{err},
	}
}

//NewUnauthorizedError create new unauthorized error status with message
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: "Unable to retrive	user information from given access token",
		Status: http.StatusUnauthorized,
		Error:  "unauthorized",
	}
}

//NewError create new error
func NewError(msg string) error {
	return errors.New(msg)
}
