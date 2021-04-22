// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*UserRegisterDefault successful operation

swagger:response userRegisterDefault
*/
type UserRegisterDefault struct {
	_statusCode int
}

// NewUserRegisterDefault creates UserRegisterDefault with default headers values
func NewUserRegisterDefault(code int) *UserRegisterDefault {
	if code <= 0 {
		code = 500
	}

	return &UserRegisterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user register default response
func (o *UserRegisterDefault) WithStatusCode(code int) *UserRegisterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user register default response
func (o *UserRegisterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *UserRegisterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
