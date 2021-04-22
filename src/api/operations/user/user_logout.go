// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/jannst/go_start/auth_service/src"
)

// UserLogoutHandlerFunc turns a function with the right signature into a user logout handler
type UserLogoutHandlerFunc func(UserLogoutParams, *src.Session) middleware.Responder

// Handle executing the request and returning a response
func (fn UserLogoutHandlerFunc) Handle(params UserLogoutParams, principal *src.Session) middleware.Responder {
	return fn(params, principal)
}

// UserLogoutHandler interface for that can handle valid user logout params
type UserLogoutHandler interface {
	Handle(UserLogoutParams, *src.Session) middleware.Responder
}

// NewUserLogout creates a new http.Handler for the user logout operation
func NewUserLogout(ctx *middleware.Context, handler UserLogoutHandler) *UserLogout {
	return &UserLogout{Context: ctx, Handler: handler}
}

/* UserLogout swagger:route POST /logout User userLogout

User logout

*/
type UserLogout struct {
	Context *middleware.Context
	Handler UserLogoutHandler
}

func (o *UserLogout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserLogoutParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *src.Session
	if uprinc != nil {
		principal = uprinc.(*src.Session) // this is really a src.Session, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}