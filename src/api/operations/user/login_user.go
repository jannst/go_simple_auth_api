// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginUserHandlerFunc turns a function with the right signature into a login user handler
type LoginUserHandlerFunc func(LoginUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginUserHandlerFunc) Handle(params LoginUserParams) middleware.Responder {
	return fn(params)
}

// LoginUserHandler interface for that can handle valid login user params
type LoginUserHandler interface {
	Handle(LoginUserParams) middleware.Responder
}

// NewLoginUser creates a new http.Handler for the login user operation
func NewLoginUser(ctx *middleware.Context, handler LoginUserHandler) *LoginUser {
	return &LoginUser{Context: ctx, Handler: handler}
}

/* LoginUser swagger:route POST /login User loginUser

User login

*/
type LoginUser struct {
	Context *middleware.Context
	Handler LoginUserHandler
}

func (o *LoginUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// LoginUserBody User login information
//
// swagger:model LoginUserBody
type LoginUserBody struct {

	// email
	// Required: true
	// Max Length: 256
	// Format: email
	Email *strfmt.Email `json:"email"`

	// password
	// Example: VerySecureLol_69
	// Required: true
	// Max Length: 256
	// Min Length: 8
	Password *string `json:"password"`
}

// Validate validates this login user body
func (o *LoginUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginUserBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"email", "body", o.Email.String(), 256); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"email", "body", "email", o.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *LoginUserBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"password", "body", *o.Password, 8); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"password", "body", *o.Password, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login user body based on context it is used
func (o *LoginUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LoginUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginUserBody) UnmarshalBinary(b []byte) error {
	var res LoginUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
