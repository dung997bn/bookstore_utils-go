package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//RestErr type
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s -causes: [%v]", e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

//NewRestError func
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

//NewRestErrorFromBytes func
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

//NewBadRequestError create new bad request status with message
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError create new not found status with message
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "bad_request",
	}
}

//NewInternalServerError create new internal server error status with message
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}

	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

//NewUnauthorizedError create new unauthorized error status with message
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: "Unable to retrive	user information from given access token",
		ErrStatus: http.StatusUnauthorized,
		ErrError:  "unauthorized",
	}
}
