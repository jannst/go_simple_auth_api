// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jannst/go_start/auth_service/src/apimodel"
)

// InfoOKCode is the HTTP code returned for type InfoOK
const InfoOKCode int = 200

/*InfoOK successful operation

swagger:response infoOK
*/
type InfoOK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.User `json:"body,omitempty"`
}

// NewInfoOK creates InfoOK with default headers values
func NewInfoOK() *InfoOK {

	return &InfoOK{}
}

// WithPayload adds the payload to the info o k response
func (o *InfoOK) WithPayload(payload *apimodel.User) *InfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the info o k response
func (o *InfoOK) SetPayload(payload *apimodel.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
