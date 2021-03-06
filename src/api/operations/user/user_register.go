// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UserRegisterHandlerFunc turns a function with the right signature into a user register handler
type UserRegisterHandlerFunc func(UserRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserRegisterHandlerFunc) Handle(params UserRegisterParams) middleware.Responder {
	return fn(params)
}

// UserRegisterHandler interface for that can handle valid user register params
type UserRegisterHandler interface {
	Handle(UserRegisterParams) middleware.Responder
}

// NewUserRegister creates a new http.Handler for the user register operation
func NewUserRegister(ctx *middleware.Context, handler UserRegisterHandler) *UserRegister {
	return &UserRegister{Context: ctx, Handler: handler}
}

/* UserRegister swagger:route POST /register User userRegister

Register new User

*/
type UserRegister struct {
	Context *middleware.Context
	Handler UserRegisterHandler
}

func (o *UserRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserRegisterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
